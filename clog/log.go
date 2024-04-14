package clog

import (
	"fmt"
	"log"
	"os"
)

type LogLevel int

const (
	Debug LogLevel = iota
	Info
	Warning
	Error
)

type Logger struct {
	*log.Logger
	logLevel LogLevel
}

func NewLogger(out *os.File, prefix string, flag int) *Logger {
	return &Logger{
		Logger: log.New(out, prefix, flag),
	}
}

func (l *Logger) SetLogLevel(level LogLevel) {
	l.logLevel = level
}

func (l *Logger) Debug(v ...interface{}) {
	if l.logLevel <= Debug {
		l.Logger.Println("[DEBUG]", fmt.Sprint(v...))
	}
}

func (l *Logger) Info(v ...interface{}) {
	if l.logLevel <= Info {
		l.Logger.Println("[INFO]", fmt.Sprint(v...))
	}
}

func (l *Logger) Warning(v ...interface{}) {
	if l.logLevel <= Warning {
		l.Logger.Println("[WARNING]", fmt.Sprint(v...))
	}
}

func (l *Logger) Error(v ...interface{}) {
	if l.logLevel <= Error {
		l.Logger.Println("[ERROR]", fmt.Sprint(v...))
	}
}

func Test(prefix string, logLevel LogLevel) {
	// 创建日志记录器
	logger := NewLogger(os.Stdout, prefix, log.Ldate|log.Ltime)

	// 设置日志级别
	logger.SetLogLevel(Info)

	// 记录日志
	logger.Debug("This is a debug message")
	logger.Info("This is an info message")
	logger.Warning("This is a warning message")
	logger.Error("This is an error message")
}
