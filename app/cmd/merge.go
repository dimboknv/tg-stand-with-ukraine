package cmd

import (
	"github.com/dimboknv/tg-stand-with-ukraine/app/store"
	"github.com/jessevdk/go-flags"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type MergeCmd struct {
	Current flags.Filename `long:"current" env:"CURRENT" description:"Current database filepath" required:"true" default:"current.db"`
	Target  flags.Filename `long:"target" env:"TARGET" description:"Target database filepath" required:"true" default:"target.db"`
	CommonOpts
}

func (cmd *MergeCmd) Execute(args []string) error {
	cmd.Logger.Info("\"merge\" command started", zap.String("current", string(cmd.Current)), zap.String("target", string(cmd.Target)))

	current, err := store.NewBoltStore(string(cmd.Current))
	if err != nil {
		return errors.Wrap(err, "can`t create bolt store")
	}
	target, err := store.NewBoltStore(string(cmd.Target))
	if err != nil {
		return errors.Wrap(err, "can`t create bolt store")
	}
	defer func() {
		if closeErr := current.Close(); err != nil {
			cmd.Logger.Warn("fail to close database", zap.Error(closeErr))
		}
		if closeErr := target.Close(); err != nil {
			cmd.Logger.Warn("fail to close database", zap.Error(closeErr))
		}
	}()

	if err := cmd.merge(&usersMerger{db: current}, target); err != nil {
		return err
	}

	cmd.Logger.Info("\"merge\" command succeeded terminated")
	return nil
}

func (cmd *MergeCmd) merge(current Merger, target store.Store) error {
	users, err := target.GetUsers()
	if err != nil {
		return err
	}
	for i := range users {
		if err := current.MergeUser(users[i]); err != nil {
			return err
		}
	}
	return nil
}

type Merger interface {
	MergeUser(user store.User) error
}

type usersMerger struct {
	db store.Store
}

func (m *usersMerger) MergeUser(user store.User) error {
	currentUser, err := m.db.GetUser(user.ID)
	if err != nil {
		if errors.Is(err, store.NotFoundError) {
			return m.db.PutUser(user)
		}
		return err
	}

	for chatID, chat := range user.Chats {
		if _, has := currentUser.Chats[chatID]; !has {
			currentUser.Chats[chatID] = chat
		}
	}

	for phone, client := range user.Clients {
		currentClient, has := currentUser.Clients[phone]
		if !has {
			currentUser.Clients[phone] = client
			continue
		}

		for url, report := range client.SentReports {
			if _, has := currentClient.SentReports[url]; !has {
				currentClient.SentReports[url] = report
			}
		}
	}

	return m.db.PutUser(currentUser)
}
