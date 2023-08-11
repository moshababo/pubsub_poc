package common

import (
	"fmt"
	"time"
)

type Logger struct {
	Level        LogLevel
	SubsystemTag string
}

type LogLevel int

const (
	LogLevelInfo LogLevel = iota
	LogLevelWarning
	LogLevelError
)

func NewLogger(level LogLevel, tag string) *Logger {
	return &Logger{Level: level, SubsystemTag: tag}
}

func (l *Logger) log(level LogLevel, msg string) {
	if level >= l.Level {
		fmt.Printf("%s [%s] %s: %s\n",
			time.Now().Format("2006-01-02 15:04:05.000"),
			levelToString(level),
			l.SubsystemTag,
			msg,
		)
	}
}

func (l *Logger) WithTag(tag string) *Logger {
	return NewLogger(l.Level, tag)
}

func (l *Logger) Info(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.log(LogLevelInfo, msg)
}

func (l *Logger) Warning(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.log(LogLevelWarning, msg)
}

func (l *Logger) Error(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.log(LogLevelError, msg)
}

func levelToString(level LogLevel) string {
	switch level {
	case LogLevelInfo:
		return "INF"
	case LogLevelWarning:
		return "WARN"
	case LogLevelError:
		return "ERR"
	default:
		return "UNKNOWN"
	}
}
