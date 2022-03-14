package store

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.etcd.io/bbolt"
)

var testDB = "/tmp/test-bot.db"

// makes new bboltdb with 2 users
func prep(t *testing.T) (db *bboltStore, teardown func()) {
	_ = os.Remove(testDB)

	b, err := NewBoltStore(testDB)
	assert.NoError(t, err)
	db = b

	uu := []User{
		{
			Clients: map[string]*Client{
				"+123456789": {
					SentReports: map[string]Report{
						"https://t.me/someuser1": {
							URL:  "https://t.me/someuser1",
							Text: "some report msg1",
						},
					},
					Session: []byte("session1"),
				},
			},
			Chats: map[int64]*Chat{
				1: {
					AuthPhone:  "+987654321",
					Navigation: 1,
				},
			},
			ID: 1,
		}, {
			Clients: map[string]*Client{
				"+123456789": {
					SentReports: map[string]Report{
						"https://t.me/someuser2": {
							URL:  "https://t.me/someuser2",
							Text: "some report msg2",
						},
					},
					Session: []byte("session2"),
				},
			},
			Chats: map[int64]*Chat{
				2: {
					AuthPhone:  "+987654321",
					Navigation: 2,
				},
			},
			ID: 2,
		},
	}

	assert.NoError(t, db.db.Update(func(tx *bbolt.Tx) error {
		for _, u := range uu {
			data, err := json.Marshal(u)
			if err != nil {
				return err
			}
			if err := tx.Bucket([]byte(usersBucket)).Put([]byte(strconv.FormatInt(u.ID, 10)), data); err != nil {
				return err
			}
		}
		return nil
	}))

	teardown = func() {
		require.NoError(t, db.Close())
		_ = os.Remove(testDB)
	}
	return db, teardown
}

func Test_bboltStore_GetUser(t *testing.T) {
	_, teardown := prep(t)
	defer teardown()
	fmt.Println("")
}
