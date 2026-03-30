package logger

import (
	"context"

	"go.uber.org/zap"
)

type Logger interface {
	Info(ctx context.Context, message string, args ...any)
	Error(ctx context.Context, message string, args ...any)
	Warn(ctx context.Context, message string, args ...any)
	Debug(ctx context.Context, message string, args ...any)
}

type logger struct {
	zap *zap.SugaredLogger
}

func NewLogger() (Logger, error) {
	zapLogger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}

	return &logger{zap: zapLogger.Sugar()}, nil
}

func (l *logger) Info(ctx context.Context, message string, args ...any) {
	if len(args) > 0 {
		l.zap.Infof(message, args...)
		return
	}
	l.zap.Info(message)
}

func (l *logger) Error(ctx context.Context, message string, args ...any) {
	if len(args) > 0 {
		l.zap.Errorf(message, args...)
		return
	}
	l.zap.Error(message)
}

func (l *logger) Warn(ctx context.Context, message string, args ...any) {
	if len(args) > 0 {
		l.zap.Warnf(message, args...)
		return
	}
	l.zap.Warn(message)
}

func (l *logger) Debug(ctx context.Context, message string, args ...any) {
	if len(args) > 0 {
		l.zap.Debugf(message, args...)
		return
	}
	l.zap.Debug(message)
}
