// Package session implements session storage.
package session

import (
	"context"
	"encoding/json"

	"github.com/go-faster/errors"

	"github.com/gotd/td/tg"
)

// Data of session.
type Data struct {
	Config    tg.Config
	DC        int
	Addr      string
	AuthKey   []byte
	AuthKeyID []byte
	Salt      int64
}

// Storage is secure persistent storage for client session.
//
// NB: Implementation security is important, attacker can abuse it not only for
// connecting as authenticated user or bot, but even decrypting previous
// messages in some situations.
type Storage interface {
	LoadSession(ctx context.Context) ([]byte, error)
	StoreSession(ctx context.Context, data []byte) error
}

// ErrNotFound means that session is not found in storage.
var ErrNotFound = errors.New("session storage: not found")

// Loader wraps Storage implementing Data (un-)marshaling.
type Loader struct {
	Storage Storage
}

type jsonData struct {
	Version int
	Data    Data
}

const latestVersion = 1

// Load loads Data from Storage.
func (l *Loader) Load(ctx context.Context) (*Data, error) {
	buf, err := l.Storage.LoadSession(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "load")
	}
	if len(buf) == 0 {
		return nil, ErrNotFound
	}

	var v jsonData
	if err := json.Unmarshal(buf, &v); err != nil {
		return nil, errors.Wrap(err, "unmarshal")
	}
	if v.Version != latestVersion {
		// HACK(ernado): backward compatibility super shenanigan.
		return nil, errors.Wrapf(ErrNotFound, "version mismatch (%d != %d)", v.Version, latestVersion)
	}
	return &v.Data, err
}

// Save saves Data to Storage.
func (l *Loader) Save(ctx context.Context, data *Data) error {
	v := jsonData{
		Version: latestVersion,
		Data:    *data,
	}
	buf, err := json.Marshal(v)
	if err != nil {
		return errors.Wrap(err, "marshal")
	}
	if err := l.Storage.StoreSession(ctx, buf); err != nil {
		return errors.Wrap(err, "store")
	}
	return nil
}
