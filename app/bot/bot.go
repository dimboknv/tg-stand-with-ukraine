package bot

import (
	"context"
	"net/http"
	"regexp"
	"time"

	"github.com/dimboknv/tg-stand-with-ukraine/app/reporter"

	"github.com/dimboknv/tg-stand-with-ukraine/app/hub"
	"github.com/dimboknv/tg-stand-with-ukraine/app/store"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Opts struct {
	DB       store.Store
	Logger   *zap.Logger
	Hub      *hub.Hub
	Reporter *reporter.Reporter
	Token    string
	Admins   []string
	Debug    bool
}

var digitsRegexp = regexp.MustCompile(`\D+`)

type Bot struct {
	hub         *hub.Hub
	db          store.Store
	bot         *tgbotapi.BotAPI
	reporter    *reporter.Reporter
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
		hub:      opts.Hub,
		db:       opts.DB,
		bot:      bot,
		admins:   admins,
		log:      opts.Logger,
		reporter: opts.Reporter,
	}, nil
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
			Chats: map[int64]*store.Chat{},
			ID:    u.SentFrom().ID,
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

func (b *Bot) Run(ctx context.Context) {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
	updates := b.bot.GetUpdatesChan(updateConfig)
	b.registerHandlers()
	go b.reporter.Run(ctx)

	b.log.Info("start listen for updates...")
	for {
		select {
		case u := <-updates:
			go func(u tgbotapi.Update) {
				defer func() {
					if r := recover(); r != nil {
						b.log.Warn("runtime panic", zap.Any("recover", r))
					}
				}()

				c, cancel := context.WithTimeout(context.Background(), 60*time.Second)
				defer cancel()
				if err := b.handleUserErrorIfNeeded(u, b.handleUpdate(c, u)); err != nil {
					b.log.Error("fail to handle user error", zap.Any("update", u), zap.Error(err))
				}
			}(u)
		case <-ctx.Done():
			b.log.Info("stop listen for updates")
			return
		}
	}
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
	user.Chats[u.FromChat().ID].Navigation = store.UserNavigation
	return b.db.PutUser(user)
}
