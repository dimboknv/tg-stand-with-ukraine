package hub

import (
	"context"
	"time"

	"github.com/dimboknv/tg-stand-with-ukraine/app/store"

	"github.com/gotd/td/telegram"
)

type action func(ctx context.Context, client *telegram.Client, user store.User, phone string) error

func chain(actions ...action) action {
	return func(ctx context.Context, client *telegram.Client, user store.User, phone string) error {
		for i := 0; i < len(actions); i++ {
			if err := actions[i](ctx, client, user, phone); err != nil {
				return err
			}
		}
		return nil
	}
}

func updateLastConnectionAt(db store.Store) action {
	return func(ctx context.Context, client *telegram.Client, user store.User, phone string) error {
		u, err := db.GetUser(user.ID)
		if err != nil {
			return err
		}
		u.Clients[phone].LastConnectionAt = time.Now()
		return db.PutUser(u)
	}
}

func updateIsAuthorized(db store.Store) action {
	return func(ctx context.Context, client *telegram.Client, user store.User, phone string) error {
		isAuthorized, err := isAuthorizedClient(ctx, client)
		if err != nil {
			return err
		}

		u, err := db.GetUser(user.ID)
		if err != nil {
			return err
		}

		u.Clients[phone].IsAuthorized = isAuthorized
		return db.PutUser(u)
	}
}

func updateLogInAt(db store.Store) action {
	return func(ctx context.Context, client *telegram.Client, user store.User, phone string) error {
		u, err := db.GetUser(user.ID)
		if err != nil {
			return err
		}
		u.Clients[phone].LogInAt = time.Now()
		return db.PutUser(u)
	}
}

func requireAuthorized(ctx context.Context, client *telegram.Client, user store.User, phone string) error {
	isAuthorized, err := isAuthorizedClient(ctx, client)
	if err != nil {
		return err
	}
	if !isAuthorized {
		return NotAuthorizedErr
	}
	return nil
}
