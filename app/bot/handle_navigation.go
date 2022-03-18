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
	digitsRegexp   = regexp.MustCompile(`\D+`)
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

// nolint:gocyclo // user has many input sources
func getMessageText(user store.User, chatID int64, u tgbotapi.Update) string {
	msg, cbq, userChat := u.Message, u.CallbackQuery, user.Chats[chatID]
	switch {
	case cbq != nil && cbq.Message != nil && cbq.Message.MessageID == userChat.ReplyMsgID:
		return cbq.Data
	case msg != nil && msg.ReplyToMessage != nil && msg.ReplyToMessage.MessageID == userChat.ReplyMsgID && msg.Contact != nil:
		return msg.Contact.PhoneNumber
	case msg != nil && msg.Contact != nil:
		return msg.Contact.PhoneNumber
	case u.Message != nil:
		return u.Message.Text
	default:
		return ""
	}
}

func parsePhone(txt string) (phone string, err error) {
	phone = "+" + digitsRegexp.ReplaceAllString(txt, "")
	//+380-44-xxx-xx-xx
	if len(phone) != 13 {
		return "", &userError{
			Err:     errors.Errorf("invalid phone number: %q", phone),
			UserMsg: "Invalid phone number. Try again /login",
		}
	}
	return phone, nil
}

func (b *Bot) phoneNavigation(ctx context.Context, user store.User, chatID int64, u tgbotapi.Update) error {
	phone, err := parsePhone(getMessageText(user, chatID, u))
	if err != nil {
		return err
	}

	txt, navigation := "Send pass code", store.CodeNavigation
	if phone == user.Phone {
		// nolint
		txt = `You are trying to sign-in with you current account and telegram does not allow to send secure code in one message to anyone.

Please, split secure code and send by 2 messages. For example, you got "12345" then send 1-th message with text "123" and 2-th message with "45". You can split code in any combination except full code.

Send 1-th part`
		navigation = store.SplitCode1Navigation
	}

	user.Chats[chatID].Navigation, user.Chats[chatID].AuthPhone = navigation, phone
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

	msg := tgbotapi.NewMessage(chatID, txt)
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(false)
	if _, err := b.bot.Send(msg); err != nil {
		return err
	}
	return nil
}

func (b *Bot) splitCode1Navigation(ctx context.Context, user store.User, chatID int64, u tgbotapi.Update) error {
	user.Chats[chatID].Navigation = store.SplitCode2Navigation
	user.Chats[chatID].AuthCode = strings.TrimSpace(getMessageText(user, chatID, u))
	user.Chats[chatID].DeleteMsgIDs = append(user.Chats[chatID].DeleteMsgIDs, u.Message.MessageID)
	if err := b.db.PutUser(user); err != nil {
		return err
	}
	return b.sendMsg(chatID, "Send 2-th part")
}

func (b *Bot) splitCode2Navigation(ctx context.Context, user store.User, chatID int64, u tgbotapi.Update) error {
	user.Chats[chatID].AuthCode += strings.TrimSpace(getMessageText(user, chatID, u))
	user.Chats[chatID].DeleteMsgIDs = append(user.Chats[chatID].DeleteMsgIDs, u.Message.MessageID)
	if err := b.db.PutUser(user); err != nil {
		return err
	}
	return b.startHubAuth(ctx, user, chatID, user.Chats[chatID].AuthCode)
}

func (b *Bot) codeNavigation(ctx context.Context, user store.User, chatID int64, u tgbotapi.Update) error {
	user.Chats[chatID].DeleteMsgIDs = append(user.Chats[chatID].DeleteMsgIDs, u.Message.MessageID)
	if err := b.db.PutUser(user); err != nil {
		return err
	}
	return b.startHubAuth(ctx, user, chatID, strings.TrimSpace(getMessageText(user, chatID, u)))
}

func (b *Bot) startHubAuth(ctx context.Context, user store.User, chatID int64, code string) error {
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
	if err := b.deleteChatMessages(chatID, user.Chats[chatID].DeleteMsgIDs...); err != nil {
		return err
	}
	msg := fmt.Sprintf("Thanks! %q is successfully sign in!", user.Chats[chatID].AuthPhone)
	user.Chats[chatID].Navigation = store.UserNavigation
	user.Chats[chatID].DeleteMsgIDs = nil
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
	pass2fa := strings.TrimSpace(getMessageText(user, chatID, u))
	user.Chats[chatID].DeleteMsgIDs = append(user.Chats[chatID].DeleteMsgIDs, u.Message.MessageID)
	if err := b.deleteChatMessages(chatID, user.Chats[chatID].DeleteMsgIDs...); err != nil {
		return err
	}
	user.Chats[chatID].Navigation = store.UserNavigation
	user.Chats[chatID].DeleteMsgIDs = nil
	if err := b.db.PutUser(user); err != nil {
		return err
	}
	if err := b.hub.AuthPass2FA(ctx, user, user.Chats[chatID].AuthPhone, pass2fa); err != nil {
		return &userError{
			Err:     errors.Wrapf(err, "can`t verify 2FA password for %q", user.Chats[chatID].AuthPhone),
			UserMsg: "Sorry, can't verify 2FA password",
		}
	}
	return b.sendMsg(chatID, fmt.Sprintf("Thanks! %q is successfully sign in!", user.Chats[chatID].AuthPhone))
}

func (b *Bot) userNavigation(ctx context.Context, user store.User, chatID int64, u tgbotapi.Update) error {
	if _, has := b.admins[u.SentFrom().UserName]; !has {
		return nil
	}

	urls := parseChanelURLs(u)

	if err := b.reporter.AddRashists(ctx, urls); err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("%d new urls added", len(urls)))
	msg.ReplyToMessageID = u.Message.MessageID
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) sharePhoneNavigation(ctx context.Context, user store.User, chatID int64, u tgbotapi.Update) error {
	phone, err := parsePhone(getMessageText(user, chatID, u))
	if err != nil {
		return err
	}

	user.Phone = phone
	user.Chats[chatID].Navigation = store.PhoneNavigation
	if err := b.db.PutUser(user); err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(chatID, "Thank")
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(false)
	if _, err := b.bot.Send(msg); err != nil {
		return err
	}

	return b.sendInlineKbWithPhone(user, chatID, phone)
}

func (b *Bot) sendInlineKbWithPhone(user store.User, chatID int64, phone string) error {
	msg := tgbotapi.NewMessage(chatID, "Send phone number")
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(phone, phone)),
	)

	resp, err := b.bot.Send(msg)
	if err != nil {
		return err
	}
	user.Chats[chatID].ReplyMsgID = resp.MessageID
	return b.db.PutUser(user)
}

func (b *Bot) deleteChatMessages(chatID int64, ids ...int) error {
	for _, id := range ids {
		if _, err := b.bot.Request(tgbotapi.NewDeleteMessage(chatID, id)); err != nil {
			return err
		}
	}
	return nil
}
