package hub

import (
	"context"

	"github.com/dimboknv/tg-stand-with-ukraine/app/store"

	"github.com/gotd/td/telegram/auth"
	"github.com/pkg/errors"
)

func (hub *Hub) AuthPhone(ctx context.Context, user store.User, phone string) error {
	client := hub.makeClient(user, phone)
	hub.runClient(phone, client)
	status, err := client.Auth().Status(ctx)
	if err != nil {
		return errors.Wrap(err, "can`t get client auth status")
	}
	if status.Authorized {
		return nil
	}

	sentCode, err := client.Auth().SendCode(ctx, phone, auth.SendCodeOptions{})
	if err != nil {
		return errors.Wrap(err, "client auth can`t send code")
	}

	hub.mu.Lock()
	hub.phoneCodeHash[phone] = sentCode.PhoneCodeHash
	hub.mu.Unlock()

	return nil
}

func (hub *Hub) AuthCode(ctx context.Context, phone, code string) (bool, error) {
	hub.mu.Lock()
	client, ok := hub.clients[phone]
	codeHash := hub.phoneCodeHash[phone]
	delete(hub.phoneCodeHash, phone)
	hub.mu.Unlock()
	if !ok {
		return false, errors.New("client is not running")
	}

	_, signInErr := client.Auth().SignIn(ctx, phone, code, codeHash)
	if errors.Is(signInErr, auth.ErrPasswordAuthNeeded) {
		return true, nil
	}

	return false, errors.Wrap(signInErr, "signIn failed")
}

func (hub *Hub) AuthPass2FA(ctx context.Context, phone, pass2fa string) error {
	hub.mu.RLock()
	client, ok := hub.clients[phone]
	hub.mu.RUnlock()
	if !ok {
		return errors.New("client is not running")
	}
	if _, err := client.Auth().Password(ctx, pass2fa); err != nil {
		return errors.Wrap(err, "invalid 2fa")
	}
	return nil
}
