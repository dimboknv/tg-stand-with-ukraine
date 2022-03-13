package cmd

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"os"
	"time"

	"github.com/dimboknv/tg-stand-with-ukraine/app/reporter"

	"github.com/dimboknv/tg-stand-with-ukraine/app/bot"
	"github.com/dimboknv/tg-stand-with-ukraine/app/hub"
	"github.com/dimboknv/tg-stand-with-ukraine/app/store"

	"github.com/jessevdk/go-flags"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Hub struct {
	AppHash     string         `long:"app_hash" env:"APP_HASH" description:"Telegram API app hash" required:"true"`
	PublicKey   flags.Filename `long:"pk" env:"PUBLIC_KEY" description:"Telegram API public key" required:"true"`
	DeviceModel string         `long:"device" env:"DEVICE" description:"Telegram API device model" default:"Dmitry Nev"`
	DCOption    struct {
		IPAddress string `long:"ip" env:"IP" description:"DC ip address" required:"true"`
		ID        int    `long:"id" env:"ID" description:"DC id" default:"2"`
		Port      int    `long:"port" env:"PORT" description:"DC port" default:"443"`
	} `group:"dc" namespace:"dc" env-namespace:"DC"`
	ClientTTL time.Duration `long:"client_ttl" env:"CLIENT_TTL" description:"A telegram API client TTL" default:"3m"`
	AppID     int           `long:"app_id" env:"APP_ID" description:"Telegram API app id" required:"true"`
}

type Reporter struct {
	Message            string        `long:"msg" env:"MESSAGE" description:"A report message" default:"The channel undermines the integrity of the Ukrainian state. Spreading fake news, misleading people. There are a lot of posts with threats against Ukrainians and Ukrainian soldiers. Block him ASAP"` // nolint
	Interval           time.Duration `long:"interval" env:"INTERVAL" description:"Interval between sending reports" default:"40m"`
	IntervalMaxReports int           `long:"max_reps" env:"INTERVAL_MAX_REPORTS" default:"25" description:"Max number of sent reports from a telegram client"` // nolint
}

type BotCmd struct {
	Reporter `group:"reporter" namespace:"reporter" env-namespace:"REPORTER"`
	Token    string `long:"token" env:"TOKEN" description:"Bot token" required:"true"`
	DB       string `long:"db" env:"DB" description:"Database filepath" required:"true" default:"bbolt.db"`
	CommonOpts
	Admins []string `short:"a" long:"admin" env:"ADMIN" env-delim:"," description:"Bot admin telegram usernames" required:"true"`
	Hub    `group:"hub" namespace:"hub" env-namespace:"HUB"`
}

// Execute gets statements list for specified merchant, entry point for "statements" command
func (cmd *BotCmd) Execute(_ []string) error {
	cmd.Logger.Info("\"bot\" command started")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		defer cancel()
		cmd.waitSigterm(ctx)
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
		Context:     ctx,
		ClientTTL:   cmd.ClientTTL,
		DB:          db,
		Logger:      cmd.Logger.Named("hub"),
		PublicKey:   publicKey,
		DeviceModel: cmd.DeviceModel,
		AppVersion:  cmd.CommonOpts.BuildInfo.Version,
		AppHash:     cmd.AppHash,
		DCOption: struct {
			IPAddress string
			ID        int
			Port      int
		}{
			IPAddress: cmd.DCOption.IPAddress,
			ID:        cmd.DCOption.ID,
			Port:      cmd.DCOption.Port,
		},
		AppID: cmd.AppID,
	})

	rep := reporter.New(reporter.Opts{
		DB:                 db,
		Hub:                h,
		Logger:             cmd.Logger.Named("rep"),
		Message:            cmd.Message,
		Interval:           cmd.Interval,
		IntervalMaxReports: cmd.IntervalMaxReports,
	})

	b, err := bot.New(bot.Opts{
		Token:    cmd.Token,
		DB:       db,
		Debug:    cmd.Debug,
		Logger:   cmd.Logger.Named("bot"),
		Hub:      h,
		Admins:   cmd.Admins,
		Reporter: rep,
	})
	if err != nil {
		return errors.Wrap(err, "can`t create bot")
	}

	b.Run(ctx)
	cmd.Logger.Info("\"bot\" command succeeded terminated")
	return nil
}
