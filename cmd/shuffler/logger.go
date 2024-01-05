// shuffler source code
// Author (c) 2023 Belousov Daniil

package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func logger_core() zapcore.Core {
	return zapcore.NewCore(
		zapcore.NewConsoleEncoder(
			zapcore.EncoderConfig{
				LevelKey:       "level",
				MessageKey:     "message",
				LineEnding:     "\n",
				EncodeLevel:    zapcore.LowercaseLevelEncoder,
				EncodeTime:     zapcore.EpochTimeEncoder,
				EncodeDuration: zapcore.SecondsDurationEncoder,
				EncodeCaller:   zapcore.ShortCallerEncoder,
			},
		),
		zapcore.AddSync(os.Stdout),
		zap.NewAtomicLevelAt(zapcore.InfoLevel),
	)
}

func logger_options() []zap.Option {
	return []zap.Option{}
}

func logger() *zap.Logger {
	return zap.New(
		logger_core(),
		logger_options()...,
	)
}
