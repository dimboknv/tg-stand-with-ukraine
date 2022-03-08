package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

// CommonOptionsCommander extends flags.Commander with SetCommon
// All commands should implement this interfaces
type CommonOptionsCommander interface {
	SetCommon(commonOpts CommonOpts)
	Execute(args []string) error
}

// BuildInfo about the executable
type BuildInfo struct {
	Version string
	Commit  string
	Date    string
}

// CommonOpts sets externally from main, shared across all commands
type CommonOpts struct {
	Logger    *zap.Logger
	BuildInfo BuildInfo
	Debug     bool
}

// SetCommon satisfies CommonOptionsCommander interface and sets common option fields
// The method called by main for each command
func (opts *CommonOpts) SetCommon(commonOpts CommonOpts) {
	opts.BuildInfo = commonOpts.BuildInfo
	opts.Debug = commonOpts.Debug
	opts.Logger = commonOpts.Logger
}

func (opts *CommonOpts) waitSigterm(ctx context.Context) error {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	defer close(sigCh)
	defer signal.Stop(sigCh)
	select {
	case <-sigCh:
		opts.Logger.Info("interrupt signal detected")
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
