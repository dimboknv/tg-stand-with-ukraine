package store

import (
	"github.com/pkg/errors"
)

type Store interface {
	GetUser(id int64) (User, error)
	PutUser(user User) error
	DeleteUser(id int64) error
	GetUsers() ([]User, error)
	GetReportURLs() ([]string, error)
	PutReportURL(url string) error
	PutReportURLs(urls []string) error
	Close() error
}

var NotFoundError = errors.New("record not found")
