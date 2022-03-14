package bot

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/dimboknv/tg-stand-with-ukraine/app/hub"

	"github.com/dimboknv/tg-stand-with-ukraine/app/store"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
)

var (
	tmeRegexp      = regexp.MustCompile(`(?m)(http://www\.|https://www\.|http://|https://)?t\.me/[^\s]*`)
	usernameRegexp = regexp.MustCompile(`(?m)@\S*`)
)

func parseChanelURLs(u tgbotapi.Update) []string {
	urls := make([]string, 0)
	for _, e := range u.Message.Entities {
		url := strings.TrimSpace(e.URL)
		if url == "" {
			continue
		}
		switch e.Type {
		case "text_link", "url":
			if tmeRegexp.MatchString(url) {
				urls = append(urls, url)
			}
		case "mention":
			urls = append(urls, url)
		}
	}
	if len(urls) > 0 {
		return urls
	}

	urls = tmeRegexp.FindAllString(u.Message.Text, -1)
	urls = append(urls, usernameRegexp.FindAllString(u.Message.Text, -1)...)

	for i := range urls {
		urls[i] = strings.TrimSpace(urls[i])
	}
	return urls
}

func (b *Bot) phoneNavigation(ctx context.Context, user store.User, chatID int64, u tgbotapi.Update) error {
	phone, m := u.Message.Text, u.Message
	switch {
	case m.ReplyToMessage != nil && m.ReplyToMessage.MessageID == user.Chats[chatID].ShareContactMsgID && m.Contact != nil:
		phone = u.Message.Contact.PhoneNumber
	case m.Contact != nil:
		phone = m.Contact.PhoneNumber
	}
	phone = "+" + digitsRegexp.ReplaceAllString(phone, "")

	//+380-44-xxx-xx-xx
	if len(phone) != 13 {
		return &userError{
			Err:     errors.Errorf("invalid phone number: %s", phone),
			UserMsg: "Invalid phone number. Try again /login",
		}
	}

	// send msg before starting auth for prevent active user client freezing with low msg id
	msg := tgbotapi.NewMessage(chatID, "Send pass code")
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(false)
	if _, err := b.bot.Send(msg); err != nil {
		return err
	}

	user.Chats[chatID].Navigation, user.Chats[chatID].AuthPhone = store.CodeNavigation, phone
	if err := b.db.PutUser(user); err != nil {
		return err
	}
	if err := b.hub.AuthPhone(ctx, user, phone); err != nil {
		if errors.Is(err, hub.AlreadyAuthorizedErr) {
			return &userError{
				Err:     err,
				UserMsg: "Already authorized",
			}
		}
		return errors.Wrapf(err, "can`t start auth for %q", phone)
	}

	return nil
}

func (b *Bot) codeNavigation(ctx context.Context, user store.User, chatID int64, u tgbotapi.Update) error {
	code := strings.TrimSpace(u.Message.Text)
	req2fa, err := b.hub.AuthCode(ctx, user, user.Chats[chatID].AuthPhone, code)
	if err != nil {
		return &userError{
			Err:     errors.Wrapf(err, "can`t verify code for %q", user.Chats[chatID].AuthPhone),
			UserMsg: "Sorry, can't verify code",
		}
	}

	if user, err = b.db.GetUser(user.ID); err != nil {
		return err
	}
	msg := "Thanks!"
	user.Chats[chatID].Navigation = store.UserNavigation
	if req2fa {
		msg = "Send 2FA code"
		user.Chats[chatID].Navigation = store.Pass2faNavigation
	}
	if err := b.db.PutUser(user); err != nil {
		return err
	}
	return b.sendMsg(chatID, msg)
}

func (b *Bot) pass2faNavigation(ctx context.Context, user store.User, chatID int64, u tgbotapi.Update) error {
	pass2fa := strings.TrimSpace(u.Message.Text)
	user.Chats[chatID].Navigation = store.UserNavigation
	if err := b.db.PutUser(user); err != nil {
		return err
	}
	if err := b.hub.AuthPass2FA(ctx, user, user.Chats[chatID].AuthPhone, pass2fa); err != nil {
		return &userError{
			Err:     errors.Wrapf(err, "can`t verify 2FA password for %q", user.Chats[chatID].AuthPhone),
			UserMsg: "Sorry, can't verify 2FA password",
		}
	}
	return b.sendMsg(chatID, "Thanks")
}

func (b *Bot) userNavigation(ctx context.Context, _ store.User, chatID int64, u tgbotapi.Update) error {
	if _, has := b.admins[u.SentFrom().UserName]; !has {
		return nil
	}

	urls := parseChanelURLs(u)

	if err := b.reporter.AddRashists(ctx, urls); err != nil {
		return err
	}

	// todo replay to the msg
	return b.sendMsg(u.FromChat().ID, fmt.Sprintf("%d urls added", len(urls)))
}

func (b *Bot) sharePhoneNavigation(ctx context.Context, user store.User, chatID int64, u tgbotapi.Update) error {
	// nolint:lll // kj
	if u.Message == nil || u.Message.ReplyToMessage == nil || u.Message.ReplyToMessage.MessageID != user.Chats[chatID].ShareContactMsgID || u.Message.Contact == nil {
		return &userError{
			Err:     errors.New("user don`t share phone number"),
			UserMsg: "You have to share phone number by \"Send phone\" button. Try again /login",
		}
	}

	user.Phone = "+" + u.Message.Contact.PhoneNumber
	user.Chats[chatID].Navigation = store.PhoneNavigation
	if err := b.db.PutUser(user); err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(chatID, "Thank")
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(false)
	if _, err := b.bot.Send(msg); err != nil {
		return err
	}
	msg = tgbotapi.NewMessage(chatID, "Send phone number")
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(user.Phone, user.Phone)))

	_, err := b.bot.Send(msg)
	return err
}
