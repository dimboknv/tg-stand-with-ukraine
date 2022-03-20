package bot

import (
	"context"

	"github.com/dimboknv/tg-stand-with-ukraine/app/store"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type handler func(ctx context.Context, user store.User, chatID int64, u tgbotapi.Update) error

func (b *Bot) registerHandlers() {
	b.navHandlers = map[store.Navigation]handler{
		store.UserNavigation:       b.userNavigation,
		store.Pass2faNavigation:    b.pass2faNavigation,
		store.CodeNavigation:       b.codeNavigation,
		store.PhoneNavigation:      b.phoneNavigation,
		store.SharePhoneNavigation: b.sharePhoneNavigation,
		store.SplitCode1Navigation: b.splitCode1Navigation,
		store.SplitCode2Navigation: b.splitCode2Navigation,
	}

	b.cmdHandlers = map[string]handler{
		store.StartCommand: b.handleStartCommand,
		store.LogInCommand: b.handleLogInCommand,
	}
}

func (b *Bot) handleUpdate(ctx context.Context, u tgbotapi.Update) error {
	chatID := u.FromChat().ID

	user, err := b.getUser(u)
	if err != nil {
		return err
	}

	handler := func(context.Context, store.User, int64, tgbotapi.Update) error { return b.sendWelcomeMsg(chatID) }
	if h, has := b.navHandlers[user.Chats[chatID].Navigation]; has {
		handler = h
	}

	if u.Message == nil {
		return handler(ctx, user, chatID, u)
	}

	// If the Message was not a command, Command() returns an empty string
	if h, has := b.cmdHandlers[u.Message.Command()]; has {
		handler = h
	}

	return handler(ctx, user, chatID, u)
}
