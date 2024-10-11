package logs

import (
	"github.com/rs/zerolog"
	"os"
)

type Logger interface {
	Debug(msg string, keysAndValues ...interface{})
	Info(msg string, keysAndValues ...interface{})
	Warn(msg string, keysAndValues ...interface{})
	Error(msg string, keysAndValues ...interface{})
	Fatal(msg string, keysAndValues ...interface{})
}

type ZeroLogger struct {
	log zerolog.Logger
}

func NewZeroLogger() Logger {
	zl := zerolog.New(os.Stdout).With().
		Timestamp().
		CallerWithSkipFrameCount(3).
		Logger()

	return &ZeroLogger{log: zl}
}

func (zl *ZeroLogger) Debug(msg string, keysAndValues ...interface{}) {
	zl.log.Debug().Fields(keysAndValues).Msg(msg)
}

func (zl *ZeroLogger) Info(msg string, keysAndValues ...interface{}) {
	zl.log.Info().Fields(keysAndValues).Msg(msg)
}

func (zl *ZeroLogger) Warn(msg string, keysAndValues ...interface{}) {
	zl.log.Warn().Fields(keysAndValues).Msg(msg)
}

func (zl *ZeroLogger) Error(msg string, keysAndValues ...interface{}) {
	zl.log.Error().Fields(keysAndValues).Msg(msg)
}

func (zl *ZeroLogger) Fatal(msg string, keysAndValues ...interface{}) {
	zl.log.Fatal().Fields(keysAndValues).Msg(msg)
}
