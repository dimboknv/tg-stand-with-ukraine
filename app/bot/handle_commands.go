package bot

import (
	"context"

	"github.com/dimboknv/tg-stand-with-ukraine/app/store"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) handleStartCommand(ctx context.Context, user store.User, u tgbotapi.Update) error {
	chatID := u.Message.Chat.ID
	user.Chats[chatID].Navigation = store.UserNavigation
	if err := b.db.PutUser(user); err != nil {
		return err
	}
	if err := b.sendWelcomeMsg(chatID); err != nil {
		return err
	}
	return b.sendMsg(chatID, "Use /login command for add telegram clients")
}

func (b *Bot) handleLoginCommand(ctx context.Context, user store.User, u tgbotapi.Update) error {
	chatID := u.Message.Chat.ID
	user.Chats[chatID].Navigation = store.PhoneNavigation
	if err := b.db.PutUser(user); err != nil {
		return err
	}
	return b.sendMsg(chatID, "Send your phone number")
}
