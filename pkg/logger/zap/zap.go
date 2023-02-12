// Package zap contains the Zap logger implementation for the Logger interface from package logger.
package zap

import (
	"fmt"

	"github.com/NEKETSKY/gg-test/pkg/logger"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InitLogger initiates new Zap logger entity. Applications should take care to
// call Sync before exiting.
func InitLogger(conf *zap.Config) (logger.Logger, error) {
	log, err := conf.Build()
	if err != nil {
		return nil, fmt.Errorf("failed to build config: %v", err)
	}
	log.Info("Logger initialized.")
	return &zapLog{log: log.Sugar()}, nil
}

// NewConfig returns new customized zap.Config. It gets required values from environmental variables "log.level"
// and "log.time.key".
func NewConfig(logLevel, timeKey string) (conf zap.Config, err error) {
	switch logLevel {
	case "info":
		conf = zap.NewProductionConfig()
	case "debug":
		conf = zap.NewDevelopmentConfig()
	default:
		return conf, fmt.Errorf("unexpected log level: %v", logLevel)
	}
	// Calling function's file name and line number are disabled to prettify are logs, because of using an interface
	// for logger steals the informative.
	conf.DisableCaller = true
	conf.EncoderConfig.TimeKey = timeKey
	conf.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	return
}

type zapLog struct {
	log *zap.SugaredLogger
}

func (z *zapLog) Infof(template string, args ...interface{}) {
	z.log.Infof(template, args...)
}

func (z *zapLog) Errorf(template string, args ...interface{}) {
	z.log.Errorf(template, args...)
}

func (z *zapLog) Fatalf(template string, args ...interface{}) {
	z.log.Fatalf(template, args...)
}
