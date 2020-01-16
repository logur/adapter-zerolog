// Package zerolog provides a Logur adapter for zerolog.
package zerolog

import (
	"context"

	"github.com/rs/zerolog"
)

// Logger is a Logur adapter for zerolog.
type Logger struct {
	logger zerolog.Logger
}

// New returns a new Logur logger.
func New(logger zerolog.Logger) *Logger {
	return &Logger{
		logger: logger,
	}
}

// Trace implements the Logur Logger interface.
func (l *Logger) Trace(msg string, fields ...map[string]interface{}) {
	// Fall back to Debug
	l.Debug(msg, fields...)
}

// Debug implements the Logur Logger interface.
func (l *Logger) Debug(msg string, fields ...map[string]interface{}) {
	event := l.logger.Debug()

	if len(fields) > 0 {
		event.Fields(fields[0])
	}

	event.Msg(msg)
}

// Info implements the Logur Logger interface.
func (l *Logger) Info(msg string, fields ...map[string]interface{}) {
	event := l.logger.Info()

	if len(fields) > 0 {
		event.Fields(fields[0])
	}

	event.Msg(msg)
}

// Warn implements the Logur Logger interface.
func (l *Logger) Warn(msg string, fields ...map[string]interface{}) {
	event := l.logger.Warn()

	if len(fields) > 0 {
		event.Fields(fields[0])
	}

	event.Msg(msg)
}

// Error implements the Logur Logger interface.
func (l *Logger) Error(msg string, fields ...map[string]interface{}) {
	event := l.logger.Error()

	if len(fields) > 0 {
		event.Fields(fields[0])
	}

	event.Msg(msg)
}

func (l *Logger) TraceContext(_ context.Context, msg string, fields ...map[string]interface{}) {
	l.Trace(msg, fields...)
}

func (l *Logger) DebugContext(_ context.Context, msg string, fields ...map[string]interface{}) {
	l.Debug(msg, fields...)
}

func (l *Logger) InfoContext(_ context.Context, msg string, fields ...map[string]interface{}) {
	l.Info(msg, fields...)
}

func (l *Logger) WarnContext(_ context.Context, msg string, fields ...map[string]interface{}) {
	l.Warn(msg, fields...)
}

func (l *Logger) ErrorContext(_ context.Context, msg string, fields ...map[string]interface{}) {
	l.Error(msg, fields...)
}
