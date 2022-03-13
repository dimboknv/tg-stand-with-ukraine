package store

import (
	"github.com/pkg/errors"
)

type Store interface {
	GetUser(id int64) (User, error)
	PutUser(user User) error
	GetUsers() ([]User, error)

	// PutReportURL(url string) error
	// PutReportURLs(urls []string) error

	GetRashist(url string) (Rashist, error)
	PutRashist(rashist Rashist) error
	GetRashists() ([]Rashist, error)
	PutRashists([]Rashist) error
	Close() error
}

var NotFoundError = errors.New("record not found")
