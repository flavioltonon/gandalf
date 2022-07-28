package zap

import (
	"github.com/flavioltonon/gandalf/common/logger"
	"go.uber.org/zap"
)

type Logger struct {
	parent *zap.Logger
}

func NewLogger() (*Logger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	return &Logger{parent: logger}, nil
}

func toZapFields(fields []logger.Field) []zap.Field {
	zapFields := make([]zap.Field, 0, len(fields))

	for _, field := range fields {
		zapFields = append(zapFields, zap.String(field.Name(), field.Value()))
	}

	return zapFields
}

func (l *Logger) Info(message string, fields ...logger.Field) {
	l.parent.Info(message, toZapFields(fields)...)
}

func (l *Logger) Error(message string, fields ...logger.Field) {
	l.parent.Error(message, toZapFields(fields)...)
}
