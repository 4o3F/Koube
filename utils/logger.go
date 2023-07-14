package utils

import (
	"github.com/rs/zerolog"
	"os"
)

var KoubeLogger zerolog.Logger

func initLogger() {
	KoubeLogger = zerolog.New(os.Stderr).With().Timestamp().Logger().Output(zerolog.ConsoleWriter{Out: os.Stderr})
	KoubeLogger.Info().Msg("Logger init complete")
}
