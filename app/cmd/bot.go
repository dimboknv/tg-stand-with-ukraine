package cmd

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"os"
	"time"

	"github.com/dimboknv/tg-stand-with-ukraine/app/bot"
	"github.com/dimboknv/tg-stand-with-ukraine/app/hub"
	"github.com/dimboknv/tg-stand-with-ukraine/app/store"

	"github.com/jessevdk/go-flags"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Hub struct {
	AppHash       string         `long:"app_hash" env:"APP_HASH" description:"Telegram API app hash" required:"true"`
	ReportMessage string         `long:"rep_msg" env:"REP_MSG" description:"A report message" default:"The channel undermines the integrity of the Ukrainian state. Spreading fake news, misleading people. There are a lot of posts with threats against Ukrainians and Ukrainian soldiers. Block him ASAP"` // nolint
	PublicKey     flags.Filename `long:"pk" env:"PUBLIC_KEY" description:"Telegram API public key" required:"true"`
	DeviceModel   string         `long:"device" env:"DEVICE" description:"Telegram API device model" default:"Dmitry Nev"`
	DCOption      struct {
		IPAddress string `long:"ip" env:"IP" description:"DC ip address" required:"true"`
		ID        int    `long:"id" env:"ID" description:"DC id" default:"2"`
		Port      int    `long:"port" env:"PORT" description:"DC port" default:"443"`
	} `group:"dc" namespace:"dc" env-namespace:"DC"`
	AppID               int           `long:"app_id" env:"APP_ID" description:"Telegram API app id" required:"true"`
	SendReportsInterval time.Duration `long:"rep_interval" env:"SEND_REPORTS_INTERVAL" description:"Interval between sending reports" default:"10m"` // nolint
	ResendReport        bool          `long:"resend_rep" env:"RESEND_REPORT" description:"Do I resend same reports?"`
}

type BotCmd struct {
	Token string `long:"token" env:"TOKEN" description:"Bot token" required:"true"`
	DB    string `long:"db" env:"DB" description:"Database filepath" required:"true" default:"bbolt.db"`
	CommonOpts
	Admins []string `short:"a" long:"admin" env:"ADMIN" description:"Bot admin telegram usernames" required:"true"`
	Hub    `group:"hub" namespace:"hub" env-namespace:"HUB"`
}

// Execute gets statements list for specified merchant, entry point for "statements" command
func (cmd *BotCmd) Execute(_ []string) error {
	cmd.Logger.Info("\"bot\" command started")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		defer cancel()
		_ = cmd.waitSigterm(ctx)
	}()

	content, err := os.ReadFile(string(cmd.PublicKey))
	if err != nil {
		return errors.Wrapf(err, "can`t read %q file", string(cmd.PublicKey))
	}
	block, _ := pem.Decode(content)
	publicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return errors.Wrap(err, "can`t parse public publicKey")
	}

	db, err := store.NewBoltStore(cmd.DB)
	if err != nil {
		return errors.Wrap(err, "can`t create bolt store")
	}
	defer func() {
		if closeErr := db.Close(); err != nil {
			cmd.Logger.Warn("fail to close database", zap.Error(closeErr))
		}
	}()

	h := hub.NewHub(hub.Opts{
		DB:                  db,
		Logger:              cmd.Logger.Named("hub"),
		PublicKey:           publicKey,
		DeviceModel:         cmd.DeviceModel,
		AppVersion:          cmd.CommonOpts.BuildInfo.Version,
		ReportMessage:       cmd.ReportMessage,
		AppHash:             cmd.AppHash,
		AppID:               cmd.AppID,
		SendReportsInterval: cmd.SendReportsInterval,
		ResendReport:        cmd.ResendReport,
		DCOption: struct {
			IPAddress string
			ID        int
			Port      int
		}{
			IPAddress: cmd.DCOption.IPAddress,
			ID:        cmd.DCOption.ID,
			Port:      cmd.DCOption.Port,
		},
	})

	b, err := bot.New(bot.Opts{
		Token:  cmd.Token,
		DB:     db,
		Debug:  true,
		Logger: cmd.Logger.Named("bot"),
		Hub:    h,
		Admins: cmd.Admins,
	})
	if err != nil {
		return errors.Wrap(err, "can`t create bot")
	}

	b.Run(ctx)
	cmd.Logger.Info("\"bot\" command succeeded terminated")
	return nil
}
