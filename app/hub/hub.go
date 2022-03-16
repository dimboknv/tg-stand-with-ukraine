package hub

import (
	"context"
	"crypto/rsa"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/dimboknv/tg-stand-with-ukraine/app/store"
	"github.com/gotd/contrib/middleware/floodwait"
	"github.com/gotd/td/bin"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/auth"
	"github.com/gotd/td/telegram/dcs"
	"github.com/gotd/td/telegram/message"
	"github.com/gotd/td/telegram/message/peer"
	"github.com/gotd/td/tg"
	"github.com/gotd/td/tgerr"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Hub struct {
	db             store.Store
	ctx            context.Context
	phoneCodeHash  map[string]string           // [userIDphone]codeHash
	clients        map[string]*telegram.Client // [userIDphone]client
	startedClients *sync.WaitGroup
	mu             *sync.RWMutex
	log            *zap.Logger
	publicKey      *rsa.PublicKey
	deviceModel    string
	appVersion     string
	appHash        string
	dcOption       struct {
		IPAddress string
		ID        int
		Port      int
	}
	clientTTL time.Duration
	appID     int
}

type Opts struct {
	Context     context.Context
	DB          store.Store
	Logger      *zap.Logger
	PublicKey   *rsa.PublicKey
	AppHash     string
	DeviceModel string
	AppVersion  string
	DCOption    struct {
		IPAddress string
		ID        int
		Port      int
	}
	ClientTTL time.Duration
	AppID     int
}

var (
	NotAuthorizedErr     = errors.New("not authorized")
	AlreadyAuthorizedErr = errors.New("already authorized")
)

func NewHub(opts Opts) *Hub {
	rand.Seed(time.Now().UnixNano())
	return &Hub{
		startedClients: &sync.WaitGroup{},
		ctx:            opts.Context,
		db:             opts.DB,
		clients:        map[string]*telegram.Client{},
		phoneCodeHash:  map[string]string{},
		mu:             &sync.RWMutex{},
		log:            opts.Logger,
		publicKey:      opts.PublicKey,
		deviceModel:    opts.DeviceModel,
		appVersion:     opts.AppVersion,
		appHash:        opts.AppHash,
		appID:          opts.AppID,
		clientTTL:      opts.ClientTTL,
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

func (h *Hub) AppID() int {
	return h.appID
}

func (h *Hub) AppHash() string {
	return h.appHash
}

func (h *Hub) SendReport(ctx context.Context, user store.User, phone, url, msg string) (store.Report, error) {
	client, err := h.getOrMakeClient(user, phone, updateLastConnectionAt(h.db), updateIsAuthorized(h.db), requireAuthorized)
	if err != nil {
		return store.Report{}, err
	}
	return h.clientSendReport(ctx, client, url, msg)
}

func (h *Hub) AuthPhone(ctx context.Context, user store.User, phone string) error {
	client, err := h.getOrMakeClient(user, phone, updateLastConnectionAt(h.db), updateIsAuthorized(h.db))
	if err != nil {
		return err
	}

	authorized, err := isAuthorizedClient(ctx, client)
	if err != nil {
		return err
	}
	if authorized {
		return AlreadyAuthorizedErr
	}

	sentCode, err := client.Auth().SendCode(ctx, phone, auth.SendCodeOptions{})
	if err != nil {
		return errors.Wrap(err, "client auth can`t send code")
	}

	h.mu.Lock()
	h.phoneCodeHash[key(user.ID, phone)] = sentCode.PhoneCodeHash
	h.mu.Unlock()
	return nil
}

func (h *Hub) AuthCode(ctx context.Context, user store.User, phone, code string) (need2fa bool, err error) {
	h.mu.Lock()
	codeHash := h.phoneCodeHash[key(user.ID, phone)]
	delete(h.phoneCodeHash, key(user.ID, phone))
	h.mu.Unlock()

	client, err := h.getClient(user.ID, phone)
	if err != nil {
		return false, err
	}

	authorized, err := isAuthorizedClient(ctx, client)
	if err != nil {
		return false, err
	}
	if authorized {
		return false, chain(updateSignInAt(h.db), updateIsAuthorized(h.db))(ctx, client, user, phone)
	}

	_, signInErr := client.Auth().SignIn(ctx, phone, code, codeHash)
	if errors.Is(signInErr, auth.ErrPasswordAuthNeeded) {
		return true, nil
	}

	return false, errors.Wrap(signInErr, "signIn failed")
}

func (h *Hub) AuthPass2FA(ctx context.Context, user store.User, phone, pass2fa string) error {
	client, err := h.getClient(user.ID, phone)
	if err != nil {
		return err
	}
	if _, err := client.Auth().Password(ctx, pass2fa); err != nil {
		return errors.Wrap(err, "invalid 2fa")
	}
	return chain(updateSignInAt(h.db), updateIsAuthorized(h.db))(ctx, client, user, phone)
}

func (h *Hub) newClient(user store.User, phone string) *telegram.Client {
	log := h.log.Named(fmt.Sprintf("%d.%s", user.ID, phone))
	return telegram.NewClient(h.appID, h.appHash, telegram.Options{
		Middlewares:    []telegram.Middleware{logFlood(log), floodwait.NewSimpleWaiter()},
		SessionStorage: NewStoreSession(user, phone, h.db),
		Logger:         log,
		Device: telegram.DeviceConfig{
			DeviceModel:    h.deviceModel,
			AppVersion:     h.appVersion,
			SystemLangCode: "en",
			LangCode:       "en",
		},
		PublicKeys: []telegram.PublicKey{{RSA: h.publicKey}},
		DCList: dcs.List{
			Options: []tg.DCOption{
				{
					ID:        h.dcOption.ID,
					IPAddress: h.dcOption.IPAddress,
					Port:      h.dcOption.Port,
				},
			},
		},
	})
}

func (h *Hub) makeClient(user store.User, phone string, onConnectActions ...action) (*telegram.Client, error) {
	client := h.newClient(user, phone)
	withErrCh := make(chan error, 1)
	h.startedClients.Add(1)

	h.mu.Lock()
	h.clients[key(user.ID, phone)] = client
	h.mu.Unlock()

	ctx, cancel := context.WithTimeout(h.ctx, h.clientTTL)

	go func() {
		defer func() {
			h.startedClients.Done()
			cancel()
			h.mu.Lock()
			delete(h.clients, key(user.ID, phone))
			h.mu.Unlock()
			defer close(withErrCh)
		}()

		h.log.Info("run telegram client", zap.String("phone", phone), zap.Duration("ttl", h.clientTTL))
		err := client.Run(ctx, func(ctx context.Context) error {
			// wait for session to be stored, because connection don`t wait session stored
			// 2022-03-16T15:32:35.582+0200 telegram/connect.go:84  Ready
			// 2022-03-16T15:32:35.687+0200 telegram/session.go:89  Data saved
			time.Sleep(500 * time.Millisecond)
			if err := chain(onConnectActions...)(ctx, client, user, phone); err != nil {
				return err
			}
			withErrCh <- nil
			<-ctx.Done()
			return ctx.Err()
		})
		withErrCh <- err
		if err != nil {
			if !errors.Is(err, context.DeadlineExceeded) {
				h.log.Error("shutdown telegram client with error", zap.String("phone", phone), zap.Error(err))
				return
			}
		}
		h.log.Info("success to shutdown telegram client", zap.String("phone", phone))
	}()

	return client, <-withErrCh
}

func (h *Hub) clientSendReport(ctx context.Context, client *telegram.Client, url, msg string) (store.Report, error) {
	api := client.API()
	p, err := message.NewSender(api).Resolve(url, peer.OnlyChannel).AsInputPeer(ctx)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled) {
			return store.Report{}, errors.Wrapf(err, "can`t resolve %q peer", url)
		}
		return store.Report{}, &ResolveURLErr{Err: err, URL: url}
	}
	ok, err := api.AccountReportPeer(ctx, &tg.AccountReportPeerRequest{
		Peer:    p,
		Reason:  &tg.InputReportReasonOther{},
		Message: msg,
	})
	if err != nil {
		return store.Report{}, errors.Wrapf(err, "fail account report peer")
	}
	if !ok {
		return store.Report{}, errors.Errorf("report is not ok")
	}

	return store.Report{
		URL:       url,
		Text:      msg,
		CreatedAt: time.Now(),
	}, nil
}

func (h *Hub) getOrMakeClient(user store.User, phone string, actions ...action) (*telegram.Client, error) {
	client, err := h.getClient(user.ID, phone)
	if err == nil {
		return client, nil
	}
	return h.makeClient(user, phone, actions...)
}

func (h *Hub) getClient(userID int64, phone string) (*telegram.Client, error) {
	h.mu.RLock()
	client, ok := h.clients[key(userID, phone)]
	h.mu.RUnlock()

	if !ok {
		return nil, errors.New("client is not running")
	}
	return client, nil
}

func (h *Hub) Run(ctx context.Context) error {
	<-ctx.Done()
	h.startedClients.Wait()
	return nil
}

func key(userID int64, phone string) string {
	return fmt.Sprintf("%d%s", userID, phone)
}

func isAuthorizedClient(ctx context.Context, client *telegram.Client) (bool, error) {
	status, err := client.Auth().Status(ctx)
	if err != nil {
		return false, errors.Wrap(err, "can`t get telegram client auth status")
	}
	return status.Authorized, nil
}

func logFlood(logger *zap.Logger) telegram.Middleware {
	return telegram.MiddlewareFunc(func(next tg.Invoker) telegram.InvokeFunc {
		return func(ctx context.Context, input bin.Encoder, output bin.Decoder) error {
			err := next.Invoke(ctx, input, output)
			if err == nil {
				return nil
			}

			d, ok := tgerr.AsFloodWait(err)
			if !ok {
				return err
			}
			logger.Info("got FLOOD_WAIT", zap.Duration("duration", d))
			return err
		}
	})
}
