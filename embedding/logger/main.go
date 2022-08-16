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
		return ""
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
	// var (
	// 	buf    bytes.Buffer
	// 	logger = log.New(&buf, "logger: ", log.Lshortfile)
	// )

	return LogExtended{
		Logger:   log.Default(), //logger,
		logLevel: LogLevelInfo,
	}
}

// helper method prints message depending on level
func (l *LogExtended) println(srcLogLvl LogLevel, prefix, msg string) {
	if l.logLevel < srcLogLvl {
		return
	}

	l.Logger.Println(prefix + msg)
}

// method Infoln
// func (l LogExtended) Infoln(msg string){
// 	l.println(l.logLevel, l.St)
// }

func main() {
	logger := NewLogExtended()
	logger.SetLogLevel(LogLevelWarning)
	log.Println(logger.Logger)
	log.Println(logger.logLevel)

	// logger.Infoln("Не должно напечататься")
	// logger.Warnln("Hello")
	// logger.Errorln("World")
	logger.Println("Debug")
}
