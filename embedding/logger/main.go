package main

import "log"

// package log
// type Logger (struct)
// method log.Println, log.Printf
//--------

type LogLevel int

// enum log levels
const (
	Info LogLevel = iota
	Warning
	Error
)

// method maps prefix to log level
func (l LogLevel) String() string {
	switch l {
	case Info:
		return ""
	case Warning:
		return "WARN"
	case Error:
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

// NewLogExtended creates an empty LogExtended
func NewLogExtended() LogExtended {
	return LogExtended{}
}

// helper method prints message depending on level
func (l *LogExtended) println(srcLogLvl LogLevel, prefix, msg string) {
	if l.logLevel < srcLogLvl {
		return
	}

	l.Logger.Println(prefix + msg)
}

func main() {
	// log.Print("Log it baby!")
	logger := NewLogExtended()
	logger.SetLogLevel(LogLevelWarning)
	logger.Infoln("Не должно напечататься")
	logger.Warnln("Hello")
	logger.Errorln("World")
	logger.Println("Debug")
}
