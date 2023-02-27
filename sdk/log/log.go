// Package log exists to make sure all code uses the journald logger
package log

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/journald"
	"runtime/debug"
)

var Logger zerolog.Logger

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	Logger = zerolog.New(journald.NewJournalDWriter())
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

func Info() *zerolog.Event {
	return Logger.Info()
}

func Fatal() *zerolog.Event {
	return Logger.Fatal()
}

func Debug() *zerolog.Event {
	return Logger.Debug()
}

func PanicHandler() {
	// Normally, each line of the printed stack trace ends up as a separate row in journald
	// We need to capture the panic, and printed the stack trace in a structured field
	if err := recover(); err != nil {
		Fatal().
			Str("panic", fmt.Sprint(err)).
			Str("stack_trace", string(debug.Stack())).
			Msg("Panic!")
	}
}
