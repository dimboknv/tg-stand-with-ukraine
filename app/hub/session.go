package hub

import (
	"context"

	"github.com/dimboknv/tg-stand-with-ukraine/app/store"

	"github.com/gotd/td/session"
)

type UserStore interface {
	GetUser(id int64) (store.User, error)
	PutUser(user store.User) error
}

type StoreSession struct {
	db    UserStore
	phone string
	user  store.User
}

func NewStoreSession(user store.User, phone string, db UserStore) *StoreSession {
	return &StoreSession{user: user, phone: phone, db: db}
}

// LoadSession loads session from store
func (s *StoreSession) LoadSession(_ context.Context) ([]byte, error) {
	if s == nil {
		return nil, session.ErrNotFound
	}

	user, err := s.db.GetUser(s.user.ID)
	if err != nil {
		return nil, err
	}

	if user.Clients == nil {
		return nil, session.ErrNotFound
	}

	cli, ok := user.Clients[s.phone]
	if !ok || len(cli.Session) == 0 {
		return nil, session.ErrNotFound
	}

	cpy := append([]byte(nil), cli.Session...)
	return cpy, nil
}

// StoreSession stores session to store
func (s *StoreSession) StoreSession(ctx context.Context, data []byte) error {
	user, err := s.db.GetUser(s.user.ID)
	if err != nil {
		return err
	}

	if user.Clients == nil {
		user.Clients = map[string]*store.Client{}
	}

	if _, ok := user.Clients[s.phone]; !ok {
		user.Clients[s.phone] = &store.Client{}
	}
	user.Clients[s.phone].Session = data

	return s.db.PutUser(user)
}
