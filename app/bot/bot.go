package bot

import (
	"context"
	"net/http"
	"regexp"
	"time"

	"github.com/dimboknv/tg-stand-with-ukraine/app/hub"
	"github.com/dimboknv/tg-stand-with-ukraine/app/store"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Opts struct {
	DB     store.Store
	Logger *zap.Logger
	Hub    *hub.Hub
	Token  string
	Admins []string
	Debug  bool
}

var digitsRegexp = regexp.MustCompile(`\D+`)

type handler func(ctx context.Context, user store.User, u tgbotapi.Update) error

type Bot struct {
	hub         *hub.Hub
	db          store.Store
	bot         *tgbotapi.BotAPI
	admins      map[string]struct{}          // [username]exist
	msgHandlers map[store.Navigation]handler // [navigation]handler
	cmdHandlers map[string]handler           // [commandName]handler
	log         *zap.Logger
}

func New(opts Opts) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(opts.Token)
	if err != nil {
		return nil, errors.Wrap(err, "can`t create telegram bot api")
	}
	_ = tgbotapi.SetLogger(newBotLogger(opts.Logger.Named("tgbotapi"), opts.Debug))
	bot.Debug = opts.Debug
	bot.Client = &http.Client{
		Timeout: time.Minute,
	}

	admins := map[string]struct{}{}
	for _, admin := range opts.Admins {
		admins[admin] = struct{}{}
	}

	return &Bot{
		hub:    opts.Hub,
		db:     opts.DB,
		bot:    bot,
		admins: admins,
		log:    opts.Logger,
	}, nil
}

func (b *Bot) getUser(u tgbotapi.Update) (store.User, error) {
	chatID := u.Message.Chat.ID
	user, err := b.db.GetUser(u.SentFrom().ID)
	if err != nil {
		if !errors.Is(err, store.NotFoundError) {
			return store.User{}, errors.Wrap(err, "can`t get user")
		}
		user = store.User{
			Chats: map[int64]*struct {
				Phone      string
				Navigation store.Navigation
			}{
				chatID: {
					Navigation: store.WelcomeNav,
				},
			},
			ID: u.SentFrom().ID,
		}
	}
	if _, has := user.Chats[chatID]; !has {
		user.Chats[chatID] = &struct {
			Phone      string
			Navigation store.Navigation
		}{
			Navigation: store.WelcomeNav,
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

func (b *Bot) handleUpdate(ctx context.Context, u tgbotapi.Update) error {
	chatID := u.Message.Chat.ID

	user, err := b.getUser(u)
	if err != nil {
		return err
	}

	handler := func(context.Context, store.User, tgbotapi.Update) error { return b.sendWelcomeMsg(chatID) }
	if h, has := b.msgHandlers[user.Chats[chatID].Navigation]; has {
		handler = h
	}

	// If the Message was not a command, Command() returns an empty string
	if h, has := b.cmdHandlers[u.Message.Command()]; has {
		handler = h
	}

	return handler(ctx, user, u)
}

func (b *Bot) Run(ctx context.Context) {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := b.bot.GetUpdatesChan(updateConfig)
	b.registerHandlers()
	go b.hub.Run(ctx)

	b.log.Info("start listen for updates...")
	for {
		select {
		case u := <-updates:
			go func(u tgbotapi.Update) {
				c, cancel := context.WithTimeout(context.Background(), 180*time.Second)
				defer cancel()
				if err := b.handleUpdate(c, u); err != nil {
					b.log.Error("fail to handle update", zap.Any("update", u), zap.Error(err))
					return
				}
			}(u)
		case <-ctx.Done():
			b.log.Info("stop listen for updates")
			return
		}
	}
}
