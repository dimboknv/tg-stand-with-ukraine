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

func (b *Bot) phoneNavigation(ctx context.Context, user store.User, u tgbotapi.Update) error {
	phone := "+" + digitsRegexp.ReplaceAllString(u.Message.Text, "")
	chatID := u.Message.Chat.ID
	user.Chats[chatID].Navigation, user.Chats[chatID].Phone = store.CodeNavigation, phone
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

	return b.sendMsg(chatID, "Send pass code")
}

func (b *Bot) codeNavigation(ctx context.Context, user store.User, u tgbotapi.Update) error {
	code := strings.TrimSpace(u.Message.Text)
	chatID := u.Message.Chat.ID

	req2fa, err := b.hub.AuthCode(ctx, user, user.Chats[chatID].Phone, code)
	if err != nil {
		return &userError{
			Err:     errors.Wrapf(err, "can`t verify code for %q", user.Chats[chatID].Phone),
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

func (b *Bot) pass2faNavigation(ctx context.Context, user store.User, u tgbotapi.Update) error {
	pass2fa := strings.TrimSpace(u.Message.Text)
	chatID := u.Message.Chat.ID
	user.Chats[chatID].Navigation = store.UserNavigation
	if err := b.db.PutUser(user); err != nil {
		return err
	}
	if err := b.hub.AuthPass2FA(ctx, user, user.Chats[chatID].Phone, pass2fa); err != nil {
		return &userError{
			Err:     errors.Wrapf(err, "can`t verify 2FA password for %q", user.Chats[chatID].Phone),
			UserMsg: "Sorry, can't verify 2FA password",
		}
	}
	return b.sendMsg(chatID, "Thanks")
}

func (b *Bot) userNavigation(ctx context.Context, _ store.User, u tgbotapi.Update) error {
	if _, has := b.admins[u.SentFrom().UserName]; !has {
		return nil
	}

	urls := parseChanelURLs(u)

	if err := b.reporter.AddRashists(ctx, urls); err != nil {
		return err
	}

	// todo replay to the msg
	return b.sendMsg(u.Message.Chat.ID, fmt.Sprintf("%d urls added", len(urls)))
}
