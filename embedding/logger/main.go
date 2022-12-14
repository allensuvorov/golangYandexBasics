package main

import (
	"log"
)

// package log
// type Logger (struct)
// method log.Println, log.Printf
//--------

type LogLevel int

// enum log levels
const (
	LogLevelInfo LogLevel = iota
	LogLevelWarning
	LogLevelError
)

// method maps prefix to log level
func (l LogLevel) String() string {
	switch l {
	case LogLevelInfo:
		return "INFO"
	case LogLevelWarning:
		return "WARN"
	case LogLevelError:
		return "ERR"
	}
	return ""
}

// func New(out io.Writer, prefix string, flag int) *Logger
// constructor for type log.Logger returns a pointer
// thus LogExtended type embedding it as a pointer
type LogExtended struct {
	*log.Logger
	logLevel LogLevel // LogLevel это enum
}

// Set log level
func (l *LogExtended) SetLogLevel(level LogLevel) {
	l.logLevel = level
}

// NewLogExtended creates an empty LogExtended
func NewLogExtended() LogExtended {
	return LogExtended{
		Logger:   log.Default(),
		logLevel: LogLevelInfo,
	}
}

// helper method prints message depending on level
func (l *LogExtended) println(srcLogLvl LogLevel, prefix, msg string) {
	if l.logLevel > srcLogLvl {
		return
	}

	l.Logger.Println(prefix + " " + msg)
}

// method Infoln(msg string);

func (l LogExtended) Infoln(msg string) {
	l.println(LogLevelInfo, LogLevelInfo.String(), msg)
}

// method Warnln(msg string);
func (l LogExtended) Warnln(msg string) {
	l.println(LogLevelWarning, LogLevelWarning.String(), msg)
}

// method Errorln(msg string)
func (l LogExtended) Errorln(msg string) {
	l.println(LogLevelError, LogLevelError.String(), msg)
}

func main() {
	logger := NewLogExtended()
	logger.SetLogLevel(LogLevelWarning)
	logger.Infoln("Не должно напечататься")
	logger.Warnln("Hello")
	logger.Errorln("World")
	logger.Println("Debug")
}
