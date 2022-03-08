package hub

import (
	"context"
	"crypto/rsa"
	"fmt"
	"sync"
	"time"

	"github.com/dimboknv/tg-stand-with-ukraine/app/store"

	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/dcs"
	"github.com/gotd/td/telegram/message"
	"github.com/gotd/td/telegram/message/peer"
	"github.com/gotd/td/tg"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Hub struct {
	db            store.Store
	ctx           context.Context
	clients       map[string]*telegram.Client
	phoneCodeHash map[string]string
	mu            *sync.RWMutex
	log           *zap.Logger
	publicKey     *rsa.PublicKey
	deviceModel   string
	appVersion    string
	reportMessage string
	appHash       string
	dcOption      struct {
		IPAddress string
		ID        int
		Port      int
	}
	sendReportsInterval time.Duration
	appID               int
	resendReport        bool
}

type Opts struct {
	DB            store.Store
	Logger        *zap.Logger
	PublicKey     *rsa.PublicKey
	DeviceModel   string
	AppVersion    string
	ReportMessage string
	AppHash       string
	DCOption      struct {
		IPAddress string
		ID        int
		Port      int
	}
	AppID               int
	SendReportsInterval time.Duration
	ResendReport        bool
}

func NewHub(opts Opts) *Hub {
	return &Hub{
		db:                  opts.DB,
		clients:             map[string]*telegram.Client{},
		phoneCodeHash:       map[string]string{},
		mu:                  &sync.RWMutex{},
		log:                 opts.Logger,
		publicKey:           opts.PublicKey,
		deviceModel:         opts.DeviceModel,
		appVersion:          opts.AppVersion,
		reportMessage:       opts.ReportMessage,
		sendReportsInterval: opts.SendReportsInterval,
		resendReport:        opts.ResendReport,
		appHash:             opts.AppHash,
		appID:               opts.AppID,
		dcOption: struct {
			IPAddress string
			ID        int
			Port      int
		}{
			ID:        opts.DCOption.ID,
			IPAddress: opts.DCOption.IPAddress,
			Port:      opts.DCOption.Port,
		},
	}
}

func (hub *Hub) AppID() int {
	return hub.appID
}

func (hub *Hub) AppHash() string {
	return hub.appHash
}

func (hub *Hub) makeClient(user store.User, phone string) *telegram.Client {
	return telegram.NewClient(hub.appID, hub.appHash, telegram.Options{
		SessionStorage: NewStoreSession(user, phone, hub.db),
		Logger:         hub.log.Named(fmt.Sprintf("%d.%s", user.ID, phone)),
		Device: telegram.DeviceConfig{
			DeviceModel:    hub.deviceModel,
			AppVersion:     hub.appVersion,
			SystemLangCode: "en",
			LangCode:       "en",
		},
		RetryInterval: 500 * time.Millisecond,
		PublicKeys:    []telegram.PublicKey{{RSA: hub.publicKey}},
		DCList: dcs.List{
			Options: []tg.DCOption{
				{
					ID:        hub.dcOption.ID,
					IPAddress: hub.dcOption.IPAddress,
					Port:      hub.dcOption.Port,
				},
			},
		},
	})
}

func (hub *Hub) sendReport(ctx context.Context, client *telegram.Client, reportURL string) (store.Report, error) {
	api := client.API()
	p, err := message.NewSender(api).Resolve(reportURL, peer.OnlyChannel).AsInputPeer(ctx)
	if err != nil {
		return store.Report{}, errors.Wrapf(err, "can`t resolve %q peer", reportURL)
	}

	ok, err := api.AccountReportPeer(ctx, &tg.AccountReportPeerRequest{
		Peer:    p,
		Reason:  &tg.InputReportReasonOther{},
		Message: hub.reportMessage,
	})
	if err != nil {
		return store.Report{}, errors.Wrapf(err, "fail to account report peer")
	}
	if !ok {
		return store.Report{}, errors.Errorf("report is not ok")
	}

	return store.Report{
		URL:  reportURL,
		Text: hub.reportMessage,
	}, nil
}

func (hub *Hub) makeAndRunClientIfNeeded(user store.User, phone string) *telegram.Client {
	hub.mu.RLock()
	client, has := hub.clients[phone]
	hub.mu.RUnlock()
	if !has {
		client = hub.makeClient(user, phone)
		// runClient will add client to hub.clients
		hub.runClient(phone, client)
	}
	return client
}

func (hub *Hub) sendReports(ctx context.Context) error {
	users, err := hub.db.GetUsers()
	if err != nil {
		return errors.Wrap(err, "can`t get users")
	}

	reportURLs, err := hub.db.GetReportURLs()
	if err != nil {
		return errors.Wrap(err, "can`t get report list")
	}

	process := func(reportURL string, user store.User, phone string) error {
		client := hub.makeAndRunClientIfNeeded(user, phone)
		rep, err := hub.sendReport(ctx, client, reportURL)
		if err != nil {
			return err
		}
		if user.Clients[phone].SentReports == nil {
			user.Clients[phone].SentReports = map[string]store.Report{}
		}
		user.Clients[phone].SentReports[reportURL] = rep
		return hub.db.PutUser(user)
	}

	for _, reportURL := range reportURLs {
		for _, user := range users {
			for phone := range user.Clients {
				if _, has := user.Clients[phone].SentReports[reportURL]; has && !hub.resendReport {
					continue
				}

				// todo mark report if account is banned
				if err := process(reportURL, user, phone); err != nil {
					hub.log.Error("failed to send report", zap.String("report", reportURL), zap.String("phone", phone), zap.Error(err))
					continue
				}
				hub.log.Info("success to send report", zap.String("report", reportURL), zap.String("phone", phone))

				// todo??
				time.Sleep(2 * time.Second)
			}
		}
	}
	return nil
}

func (hub *Hub) runClient(phone string, client *telegram.Client) {
	isStarted := make(chan struct{})
	defer close(isStarted)

	go func() {
		defer func() {
			hub.mu.Lock()
			delete(hub.clients, phone)
			hub.mu.Unlock()
		}()

		hub.log.Info("running telegram client...", zap.String("phone", phone))

		err := client.Run(hub.ctx, func(ctx context.Context) error {
			hub.mu.Lock()
			hub.clients[phone] = client
			hub.mu.Unlock()

			hub.log.Info("telegram client is started", zap.String("phone", phone))
			isStarted <- struct{}{}
			<-ctx.Done()
			return nil
		})
		if err != nil {
			hub.log.Error("shutdown telegram client with error", zap.String("phone", phone), zap.Error(err))
			return
		}
		hub.log.Info("success to shutdown telegram client", zap.String("phone", phone))
	}()
	<-isStarted
}

func (hub *Hub) Run(ctx context.Context) {
	ticker := time.NewTicker(hub.sendReportsInterval)
	defer ticker.Stop()
	hub.ctx = ctx

	for {
		select {
		case <-ticker.C:
			go func() {
				hub.log.Info("start reporting...")
				c, cancel := context.WithTimeout(ctx, 5*time.Minute)
				defer cancel()

				if err := hub.sendReports(c); err != nil {
					hub.log.Error("reporting failed", zap.Error(err))
					return
				}
				hub.log.Info("reporting is success")
			}()
		case <-ctx.Done():
			return
		}
	}
}
