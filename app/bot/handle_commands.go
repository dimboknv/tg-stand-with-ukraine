package bot

import (
	"context"

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
	return b.sendMsg(chatID, "Use /login command for add telegram clients")
}

func (b *Bot) handleLoginCommand(ctx context.Context, user store.User, chatID int64, u tgbotapi.Update) error {
	if user.Phone != "" {
		user.Chats[chatID].Navigation = store.PhoneNavigation
		if err := b.db.PutUser(user); err != nil {
			return err
		}
		return b.sendMsg(chatID, "Send phone number")
	}

	user.Chats[chatID].Navigation = store.SharePhoneNavigation
	if err := b.db.PutUser(user); err != nil {
		return err
	}

	// getting user phone number
	msg := tgbotapi.NewMessage(chatID, "Before SignIn you have to share phone number")
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButtonContact("\xF0\x9F\x93\x9E Send phone"),
		),
	)
	resp, err := b.bot.Send(msg)
	if err != nil {
		return err
	}
	user.Chats[chatID].ShareContactMsgID = resp.MessageID
	return b.db.PutUser(user)
}
