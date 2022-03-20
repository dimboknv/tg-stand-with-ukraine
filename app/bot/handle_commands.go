package bot

import (
	"context"
	"fmt"

	"github.com/dimboknv/tg-stand-with-ukraine/app/store"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) handleStartCommand(ctx context.Context, user store.User, chatID int64, u tgbotapi.Update) error {
	user.Chats[chatID].Navigation = store.UserNavigation
	if err := b.db.PutUser(user); err != nil {
		return err
	}
	if err := b.sendWelcomeMsg(chatID); err != nil {
		return err
	}
	return b.sendTextMsg(chatID, fmt.Sprintf("Use /%s command for add telegram clients", store.LogInCommand))
}

func (b *Bot) handleLogInCommand(ctx context.Context, user store.User, chatID int64, u tgbotapi.Update) error {
	if user.Phone != "" {
		user.Chats[chatID].Navigation = store.PhoneNavigation
		return b.sendInlineKbWithPhone(user, chatID, user.Phone)
	}

	user.Chats[chatID].Navigation = store.SharePhoneNavigation
	if err := b.db.PutUser(user); err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(chatID, "Before log in you have to share phone number")
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButtonContact("\xF0\x9F\x93\x9E Send phone"),
		),
	)
	resp, err := b.bot.Send(msg)
	if err != nil {
		return err
	}
	user.Chats[chatID].ReplyMsgID = resp.MessageID
	return b.db.PutUser(user)
}
