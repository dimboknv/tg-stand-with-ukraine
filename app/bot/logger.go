package bot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

type logger struct {
	log func(format string, v ...interface{})
}

func (l *logger) Println(v ...interface{}) {
	l.log(fmt.Sprintln(v...))
}

func (l *logger) Printf(format string, v ...interface{}) {
	l.log(format, v...)
}

// newBotLogger do logging only in debug mode
func newBotLogger(log *zap.Logger, debug bool) tgbotapi.BotLogger {
	l := &logger{
		log: func(format string, v ...interface{}) {},
	}

	if debug {
		l.log = func(format string, v ...interface{}) {
			log.Debug(fmt.Sprintf(format, v...))
		}
	}

	return l
}
