package zap

import "go.uber.org/zap"

type Logger struct {
	parent *zap.Logger
}

func NewLogger() (*Logger, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}

	return &Logger{parent: logger}, nil
}

func (l *Logger) Info(message string) { l.parent.Info(message) }
