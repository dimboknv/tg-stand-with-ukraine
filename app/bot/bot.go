package bot

import (
	"context"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/dimboknv/tg-stand-with-ukraine/app/hub"
	"github.com/dimboknv/tg-stand-with-ukraine/app/reporter"
	"github.com/dimboknv/tg-stand-with-ukraine/app/store"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

type Opts struct {
	DB         store.Store
	Logger     *zap.Logger
	Hub        *hub.Hub
	Reporter   *reporter.Reporter
	Token      string
	CertFile   string
	KeyFile    string
	WebhookURL string
	Pattern    string
	Address    string
	Admins     []string
	Debug      bool
}

type Bot struct {
	db          store.Store
	hub         *hub.Hub
	bot         *tgbotapi.BotAPI
	reporter    *reporter.Reporter
	admins      map[string]struct{}          // [username]exist
	navHandlers map[store.Navigation]handler // [navigation]handler
	cmdHandlers map[string]handler           // [commandName]handler
	log         *zap.Logger
	pattern     string
	certFile    string
	keyFile     string
	webhookURL  string
	address     string
}

func New(opts Opts) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPIWithClient(opts.Token, tgbotapi.APIEndpoint, &http.Client{
		Timeout: 60 * time.Second,
	})
	if err != nil {
		return nil, errors.Wrap(err, "can`t create telegram bot api")
	}

	_ = tgbotapi.SetLogger(newBotLogger(opts.Logger.Named("tgbotapi"), opts.Debug))
	bot.Debug = opts.Debug

	admins := map[string]struct{}{}
	for _, admin := range opts.Admins {
		admins[admin] = struct{}{}
	}

	return &Bot{
		hub:        opts.Hub,
		db:         opts.DB,
		bot:        bot,
		reporter:   opts.Reporter,
		admins:     admins,
		certFile:   opts.CertFile,
		keyFile:    opts.KeyFile,
		webhookURL: opts.WebhookURL,
		address:    opts.Address,
		log:        opts.Logger,
		pattern:    opts.Pattern,
	}, nil
}

// Run all deps and bot. Returns an error if occur
// nolint:gocyclo // running all deps is complexity
func (b *Bot) Run(ctx context.Context) error {
	b.registerNavigationHandlers()
	updates, srv, err := b.updates()
	if err != nil {
		return err
	}

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		switch {
		case srv == nil:
			return nil
		case b.certFile != "" && b.keyFile != "":
			return srv.ListenAndServeTLS(b.certFile, b.keyFile)
		default:
			return srv.ListenAndServe()
		}
	})
	g.Go(func() error {
		<-ctx.Done()
		if srv == nil {
			return nil
		}
		c, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()
		return srv.Shutdown(c)
	})
	g.Go(func() error {
		return b.reporter.Run(ctx)
	})
	g.Go(func() error {
		return b.hub.Run(ctx)
	})
	g.Go(func() error {
		b.log.Info("start listen for updates...")
		defer b.log.Info("stop listen for updates")
		defer b.bot.StopReceivingUpdates()

		for {
			select {
			case u := <-updates:
				go func(u tgbotapi.Update) {
					defer func() {
						if r := recover(); r != nil {
							b.log.Warn("runtime panic", zap.Any("recover", r))
						}
					}()

					c, cancel := context.WithTimeout(ctx, 60*time.Second)
					defer cancel()
					if err := b.handleUserErrorIfNeeded(u, b.handleUpdate(c, u)); err != nil {
						b.log.Error("fail to handle user error", zap.Any("update", u), zap.Error(err))
					}
				}(u)
			case <-ctx.Done():
				return nil
			}
		}
	})
	return g.Wait()
}

func (b *Bot) pullUpdates() (tgbotapi.UpdatesChannel, error) {
	if _, err := b.bot.Request(tgbotapi.DeleteWebhookConfig{}); err != nil {
		return nil, err
	}
	whInfo, err := b.bot.GetWebhookInfo()
	if err != nil {
		return nil, err
	}
	if whInfo.IsSet() {
		return nil, errors.Errorf("can`t delete webhook")
	}

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
	return b.bot.GetUpdatesChan(updateConfig), nil
}

func (b *Bot) webhookUpdates() (tgbotapi.UpdatesChannel, *http.Server, error) {
	whURL, err := url.Parse(b.webhookURL)
	if err != nil {
		return nil, nil, err
	}
	whURL.Path = path.Join(whURL.Path, b.bot.Token)

	whCfg, err := tgbotapi.NewWebhook(whURL.String())
	if err != nil {
		return nil, nil, err
	}
	if b.certFile != "" && b.keyFile != "" {
		whCfg, err = tgbotapi.NewWebhookWithCert(whURL.String(), tgbotapi.FilePath(b.certFile))
		if err != nil {
			return nil, nil, err
		}
	}

	if _, err = b.bot.Request(whCfg); err != nil {
		return nil, nil, err
	}

	whInfo, err := b.bot.GetWebhookInfo()
	if err != nil {
		return nil, nil, err
	}
	if !whInfo.IsSet() {
		return nil, nil, errors.Errorf("can`t set webhook")
	}

	return b.bot.ListenForWebhook(path.Join(b.pattern, b.bot.Token)), &http.Server{
		Addr:    b.address,
		Handler: http.DefaultServeMux,
	}, nil
}

func (b *Bot) updates() (tgbotapi.UpdatesChannel, *http.Server, error) {
	if b.webhookURL != "" {
		return b.webhookUpdates()
	}

	updates, err := b.pullUpdates()
	if err != nil {
		return nil, nil, err
	}
	return updates, nil, nil
}

func (b *Bot) getUser(u tgbotapi.Update) (store.User, error) {
	if u.SentFrom() == nil {
		return store.User{}, errors.Errorf("can`t get message sender")
	}

	chatID := u.FromChat().ID
	user, err := b.db.GetUser(u.SentFrom().ID)
	if err != nil {
		if !errors.Is(err, store.NotFoundError) {
			return store.User{}, errors.Wrap(err, "can`t get user")
		}
		user = store.User{
			Chats:   map[int64]*store.Chat{},
			ID:      u.SentFrom().ID,
			Clients: map[string]*store.Client{},
		}
	}
	if _, has := user.Chats[chatID]; !has {
		user.Chats[chatID] = &store.Chat{
			Navigation: store.UserNavigation,
			ID:         chatID,
		}
	}
	return user, nil
}

func (b *Bot) sendMsg(chatID int64, msg string) error {
	m := tgbotapi.NewMessage(chatID, msg)
	m.ParseMode = tgbotapi.ModeMarkdown
	if _, err := b.bot.Send(m); err != nil {
		return errors.Wrap(err, "fail to send bot message")
	}
	return nil
}

func (b *Bot) sendWelcomeMsg(chatID int64) error {
	return b.sendMsg(chatID, "Welcome to reporter bot. You need to give me access to your telegram account. Stand with Ukraine!")
}

func (b *Bot) handleUserErrorIfNeeded(u tgbotapi.Update, maybeUserErr error) error {
	if maybeUserErr == nil {
		return nil
	}

	chatID := u.FromChat().ID
	txt := "Ops, something goes wrong :("

	if userErr := (&userError{}); errors.As(maybeUserErr, &userErr) {
		txt = userErr.UserMsg
		b.log.Info("get user input error", zap.Error(userErr))
	} else {
		b.log.Error("fail to handle update", zap.Any("update", u), zap.Error(maybeUserErr))
	}

	msg := tgbotapi.NewMessage(chatID, txt)
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(false)
	if _, err := b.bot.Send(msg); err != nil {
		return err
	}
	user, err := b.getUser(u)
	if err != nil {
		return err
	}
	user.Chats[u.FromChat().ID].DeleteMsgIDs = nil
	user.Chats[u.FromChat().ID].Navigation = store.UserNavigation
	return b.db.PutUser(user)
}
