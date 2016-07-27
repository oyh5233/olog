package olog

import (
	"fmt"
	"log"
	"os"
)

const (
	TVerbose = iota
	TInfo
	TWarn
	TError
)

var prefix = map[Level]string{TVerbose: "VERBOSE: ", TInfo: "INFO        : ", TWarn: "WARNING: ", TError: "ERROR      : "}

type Level int

type Logger struct {
	level   Level
	loggers map[Level]log.Logger
}

var _logger *Logger = New(TVerbose)

func New(level Level) *Logger {
	var logger Logger
	logger.level = level
	logger.loggers = make(map[Level]log.Logger, TError)
	for i := TVerbose; i <= TError; i++ {
		log := log.New(os.Stdout, prefix[Level(i)], log.Lshortfile|log.Ltime)
		logger.loggers[Level(i)] = *log
	}
	return &logger
}

func Verbose(format string, v ...interface{}) {
	_logger.Log(TVerbose, format, v...)
}

func Info(format string, v ...interface{}) {
	_logger.Log(TInfo, format, v...)
}

func Warn(format string, v ...interface{}) {
	_logger.Log(TWarn, format, v...)
}

func Error(format string, v ...interface{}) {
	_logger.Log(TError, format, v...)
}

func (logger *Logger) Log(level Level, format string, v ...interface{}) {
	if level >= logger.level {
		log := logger.loggers[level]
		log.Output(3, fmt.Sprintf(format, v...))
	}
}

func (logger *Logger) Verbose(format string, v ...interface{}) {
	logger.Log(TVerbose, format, v...)
}

func (logger *Logger) Info(format string, v ...interface{}) {
	logger.Log(TInfo, format, v...)
}

func (logger *Logger) Warn(format string, v ...interface{}) {
	logger.Log(TWarn, format, v...)
}

func (logger *Logger) Error(format string, v ...interface{}) {
	logger.Log(TError, format, v...)
}
