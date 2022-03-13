package hub

import (
	"context"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/tg"
	"github.com/gotd/td/tgerr"
	"go.uber.org/zap"
)

func logFlood(logger *zap.Logger) telegram.Middleware {
	return telegram.MiddlewareFunc(func(next tg.Invoker) telegram.InvokeFunc {
		return func(ctx context.Context, input bin.Encoder, output bin.Decoder) error {
			err := next.Invoke(ctx, input, output)
			if err == nil {
				return nil
			}

			d, ok := tgerr.AsFloodWait(err)
			if !ok {
				return err
			}
			logger.Info("got FLOOD_WAIT", zap.Duration("duration", d))
			return err
		}
	})
}
