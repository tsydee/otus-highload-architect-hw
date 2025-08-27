package logger

import (
	"fmt"
	"os"

	"github.com/tsydim/otus-highload-architect-hw/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger = *zap.SugaredLogger

func New(cfg *config.Config) (Logger, error) {
	var logLevel zapcore.Level
	err := logLevel.UnmarshalText([]byte(cfg.Logger.LogLevel))
	if err != nil {
		return nil, fmt.Errorf("new logger: %w", err)
	}
	l := zap.New(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(zapcore.EncoderConfig{
				TimeKey:        "ts",
				LevelKey:       "level",
				NameKey:        "logger",
				MessageKey:     "message",
				StacktraceKey:  "stacktrace",
				CallerKey:      "caller",
				LineEnding:     zapcore.DefaultLineEnding,
				EncodeLevel:    zapcore.LowercaseLevelEncoder,
				EncodeTime:     zapcore.ISO8601TimeEncoder,
				EncodeDuration: zapcore.SecondsDurationEncoder,
				EncodeCaller:   zapcore.ShortCallerEncoder,
			}),
			zapcore.AddSync(os.Stdout),
			zap.NewAtomicLevelAt(logLevel),
		),
		zap.AddCaller(),
	)

	zap.ReplaceGlobals(l)
	return l.Sugar(), nil
}
