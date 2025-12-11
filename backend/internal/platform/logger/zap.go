package logger

import (
	"fmt"
	"sync"

	"go.uber.org/zap"
)

type Logger struct {
	zap *zap.Logger
	mu  sync.RWMutex
}

func New(environment string) (*Logger, error) {
	var zapLogger *zap.Logger
	var err error

	switch environment {
	case "development":
		zapLogger, err = zap.NewDevelopment()
	case "production":
		zapLogger, err = zap.NewProduction()
	default:
		return nil, fmt.Errorf("Unknown environment: %s", environment)
	}

	if err != nil {
		return nil, err
	}

	return &Logger{
		zap: zapLogger,
	}, nil
}

func (l *Logger) Info(msg string, fields ...zap.Field) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	if l.zap != nil {
		l.zap.Info(msg, fields...)
	}
}

func (l *Logger) Warn(msg string, fields ...zap.Field) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	if l.zap != nil {
		l.zap.Warn(msg, fields...)
	}
}

func (l *Logger) Error(msg string, fields ...zap.Field) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	if l.zap != nil {
		l.zap.Error(msg, fields...)
	}
}

func (l *Logger) Fatal(msg string, fields ...zap.Field) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	if l.zap != nil {
		l.zap.Fatal(msg, fields...)
	}
}

func (l *Logger) Debug(msg string, fields ...zap.Field) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	if l.zap != nil {
		l.zap.Debug(msg, fields...)
	}
}

func (l *Logger) Sync() error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.zap != nil {
		return l.zap.Sync()
	}
	return nil
}
