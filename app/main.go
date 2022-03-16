package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dimboknv/tg-stand-with-ukraine/app/cmd"
	"github.com/jessevdk/go-flags"
	"go.uber.org/zap"
)

// Opts with all cli commands and flags
type Opts struct {
	MergeCmd cmd.MergeCmd `command:"merge" description:"Merge target database into current"`
	BotCmd   cmd.BotCmd   `command:"bot" description:"Run bot"`
	Debug    bool         `long:"debug"  env:"DEBUG" description:"Is debug mode?"`
}

var (
	version = "0.0.1"
	commit  = "unknown"
	date    = time.Now().Format(time.RFC3339)
)

func main() {
	var opts Opts
	p := flags.NewParser(&opts, flags.HelpFlag|flags.PassDoubleDash)
	p.CommandHandler = func(command flags.Commander, args []string) error {
		logger := newLogger(opts.Debug)
		defer func() {
			_ = logger.Sync()
		}()

		c := command.(cmd.CommonOptionsCommander)
		c.SetCommon(cmd.CommonOpts{
			Debug:  opts.Debug,
			Logger: logger,
			BuildInfo: cmd.BuildInfo{
				Version: version,
				Commit:  commit,
				Date:    date,
			},
		})

		err := command.Execute(args)
		if err != nil {
			logger.Error("command failed", zap.String("command", p.Active.Name), zap.Error(err))
		}
		return err
	}

	if _, err := p.Parse(); err != nil {
		// internal flags.Error error like 'option `-o1, --option1' uses the same long name as option `-o2, --option1'
		// wouldn't be printed by flags.Default
		w, code := os.Stderr, 1
		if flagsErr := (&flags.Error{}); errors.As(err, &flagsErr) && flagsErr.Type == flags.ErrHelp {
			w, code = os.Stdout, 0
		}
		_, _ = fmt.Fprintln(w, err)
		os.Exit(code)
	}
}

func newLogger(debug bool) *zap.Logger {
	level := zap.NewAtomicLevelAt(zap.InfoLevel)
	if debug {
		level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}
	cfg := zap.Config{
		Level:            level,
		Development:      debug,
		Encoding:         "console",
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
	logger, _ := cfg.Build()
	return logger
}
