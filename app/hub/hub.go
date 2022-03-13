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
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/auth"
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
	clients       map[string]*telegram.Client // [userIDphone]client
	phoneCodeHash map[string]string           // [userIDphone]codeHash
	mu            *sync.RWMutex
	log           *zap.Logger
	publicKey     *rsa.PublicKey
	deviceModel   string
	appVersion    string
	appHash       string
	dcOption      struct {
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
		ctx:           opts.Context,
		db:            opts.DB,
		clients:       map[string]*telegram.Client{},
		phoneCodeHash: map[string]string{},
		mu:            &sync.RWMutex{},
		log:           opts.Logger,
		publicKey:     opts.PublicKey,
		deviceModel:   opts.DeviceModel,
		appVersion:    opts.AppVersion,
		appHash:       opts.AppHash,
		appID:         opts.AppID,
		clientTTL:     opts.ClientTTL,
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
	client, err := h.getOrMakeClient(user, phone, checkAuthorization)
	if err != nil {
		return store.Report{}, err
	}
	return h.clientSendReport(ctx, client, url, msg)
}

func (h *Hub) AuthPhone(ctx context.Context, user store.User, phone string) error {
	client, err := h.getOrMakeClient(user, phone, noop)
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
		return false, nil
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
	return nil
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

func (h *Hub) makeClient(user store.User, phone string, fn prepareFn) (*telegram.Client, error) {
	client := h.newClient(user, phone)
	withErrCh := make(chan error, 1)

	h.mu.Lock()
	h.clients[key(user.ID, phone)] = client
	h.mu.Unlock()

	ctx, cancel := context.WithTimeout(h.ctx, h.clientTTL)

	go func() {
		defer func() {
			cancel()
			h.mu.Lock()
			delete(h.clients, key(user.ID, phone))
			h.mu.Unlock()
			defer close(withErrCh)
		}()

		h.log.Info("run telegram client", zap.String("phone", phone), zap.Duration("ttl", h.clientTTL))
		err := client.Run(ctx, func(ctx context.Context) error {
			// todo save user state
			if err := fn(ctx, client); err != nil {
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

func (h *Hub) getOrMakeClient(user store.User, phone string, fn prepareFn) (*telegram.Client, error) {
	client, err := h.getClient(user.ID, phone)
	if err == nil {
		return client, nil
	}
	return h.makeClient(user, phone, fn)
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

type prepareFn func(ctx context.Context, client *telegram.Client) error

func noop(context.Context, *telegram.Client) error { return nil }

func checkAuthorization(ctx context.Context, client *telegram.Client) error {
	isAuthorized, err := isAuthorizedClient(ctx, client)
	if err != nil {
		return err
	}
	if !isAuthorized {
		return NotAuthorizedErr
	}
	return nil
}
