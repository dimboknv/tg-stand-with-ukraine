package store

import (
	"encoding/json"
	"strconv"

	"github.com/pkg/errors"
	"go.etcd.io/bbolt"
)

const (
	usersBucket    = "users"
	reportsBucket  = "reports"
	rashistsBucket = "rashists"
)

type bboltStore struct {
	db *bbolt.DB
}

func NewBoltStore(filename string) (*bboltStore, error) {
	db, err := bbolt.Open(filename, 0o666, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "can`t open %q file", filename)
	}
	store := &bboltStore{db: db}
	if err := store.init(); err != nil {
		return nil, err
	}
	return store, nil
}

func (db *bboltStore) GetRashist(url string) (Rashist, error) {
	var rashist Rashist
	err := db.db.View(func(tx *bbolt.Tx) (err error) {
		bkt := tx.Bucket([]byte(usersBucket))
		return db.load(bkt, []byte(url), &rashist)
	})
	if err != nil {
		return Rashist{}, err
	}
	return rashist, nil
}

func (db *bboltStore) PutRashist(rashist Rashist) error {
	return db.db.Update(func(tx *bbolt.Tx) error {
		bkt := tx.Bucket([]byte(rashistsBucket))
		return db.save(bkt, []byte(rashist.URL), rashist)
	})
}

func (db *bboltStore) GetRashists() ([]Rashist, error) {
	res := make([]Rashist, 0)
	err := db.db.View(func(tx *bbolt.Tx) (err error) {
		bkt := tx.Bucket([]byte(rashistsBucket))
		cursor := bkt.Cursor()
		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			var rashist Rashist
			if err := json.Unmarshal(v, &rashist); err != nil {
				return errors.Wrap(err, "failed to unmarshal")
			}
			res = append(res, rashist)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (db *bboltStore) PutRashists(rashists []Rashist) error {
	return db.db.Update(func(tx *bbolt.Tx) error {
		bkt := tx.Bucket([]byte(rashistsBucket))
		for _, rashist := range rashists {
			if err := db.save(bkt, []byte(rashist.URL), rashist); err != nil {
				return err
			}
		}
		return nil
	})
}

func (db *bboltStore) GetReportURLs() ([]string, error) {
	res := make([]string, 0)
	err := db.db.View(func(tx *bbolt.Tx) (err error) {
		bkt := tx.Bucket([]byte(reportsBucket))
		cursor := bkt.Cursor()
		for k, _ := cursor.First(); k != nil; k, _ = cursor.Next() {
			res = append(res, string(k))
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (db *bboltStore) PutReportURL(url string) error {
	return db.db.Update(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte(reportsBucket)).Put([]byte(url), nil)
	})
}

func (db *bboltStore) PutReportURLs(urls []string) error {
	return db.db.Update(func(tx *bbolt.Tx) error {
		for _, url := range urls {
			if err := tx.Bucket([]byte(reportsBucket)).Put([]byte(url), nil); err != nil {
				return err
			}
		}
		return nil
	})
}

func (db *bboltStore) GetUsers() ([]User, error) {
	res := make([]User, 0)
	err := db.db.View(func(tx *bbolt.Tx) (err error) {
		bkt := tx.Bucket([]byte(usersBucket))
		cursor := bkt.Cursor()
		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			var user User
			if err := json.Unmarshal(v, &user); err != nil {
				return errors.Wrap(err, "failed to unmarshal")
			}
			res = append(res, user)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (db *bboltStore) GetUser(id int64) (User, error) {
	var user User
	err := db.db.View(func(tx *bbolt.Tx) (err error) {
		bkt := tx.Bucket([]byte(usersBucket))
		return db.load(bkt, []byte(strconv.FormatInt(id, 10)), &user)
	})
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (db *bboltStore) PutUser(user User) error {
	return db.db.Update(func(tx *bbolt.Tx) error {
		bkt := tx.Bucket([]byte(usersBucket))
		return db.save(bkt, []byte(strconv.FormatInt(user.ID, 10)), user)
	})
}

func (db *bboltStore) DeleteUser(id int64) error {
	return db.db.Update(func(tx *bbolt.Tx) error {
		bkt := tx.Bucket([]byte(usersBucket))
		return bkt.Delete([]byte(strconv.FormatInt(id, 10)))
	})
}

// save marshaled value to key for bucket. Should Run in update tx
func (db *bboltStore) save(bkt *bbolt.Bucket, key []byte, value interface{}) (err error) {
	if value == nil {
		return errors.Errorf("can't save nil value for %s", key)
	}
	jdata, jerr := json.Marshal(value)
	if jerr != nil {
		return errors.Wrap(jerr, "can't marshal item")
	}
	if err = bkt.Put(key, jdata); err != nil {
		return errors.Wrapf(err, "failed to save key %s", key)
	}
	return nil
}

// load and unmarshal json value by key from bucket. Should Run in view tx
func (db *bboltStore) load(bkt *bbolt.Bucket, key []byte, res interface{}) error {
	value := bkt.Get(key)
	if value == nil {
		return NotFoundError
	}

	if err := json.Unmarshal(value, &res); err != nil {
		return errors.Wrap(err, "failed to unmarshal")
	}
	return nil
}

func (db *bboltStore) init() error {
	return db.db.Update(func(tx *bbolt.Tx) (err error) {
		bucketNames := []string{usersBucket, reportsBucket, rashistsBucket}
		for _, bucketName := range bucketNames {
			if _, err = tx.CreateBucketIfNotExists([]byte(bucketName)); err != nil {
				return errors.Wrapf(err, "can`t create %q bucket", bucketName)
			}
		}
		return nil
	})
}

func (db *bboltStore) Close() error {
	return db.db.Close()
}
