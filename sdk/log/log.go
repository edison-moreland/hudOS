package log

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/journald"
)

var Logger zerolog.Logger

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	Logger = zerolog.New(journald.NewJournalDWriter())
}

func Log() *zerolog.Event {
	return Logger.Log()
}

func Error() *zerolog.Event {
	return Logger.Error()
}

func Panic() *zerolog.Event {
	return Logger.Panic()
}

func Warn() *zerolog.Event {
	return Logger.Warn()
}
