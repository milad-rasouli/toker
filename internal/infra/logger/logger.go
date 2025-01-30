package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

func New(cfg Config) (*zap.Logger, func() error) {
	var lvl zapcore.Level
	if err := lvl.Set(cfg.Level); err != nil {
		log.Printf("failed parse log level %s: %s\n", cfg.Level, err)
		lvl = zapcore.WarnLevel
	}

	zapCfg := zap.NewDevelopmentConfig()
	zapCfg.Level.SetLevel(lvl)

	logger, err := zapCfg.Build()
	if err != nil {
		log.Fatalf("failed to create logger %s\n", err)
	}

	return logger, func() error { return logger.Sync() }
}
