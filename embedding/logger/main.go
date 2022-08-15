package main

import "log"

// package log
// type Logger (struct)
// method log.Println, log.Printf
//--------

// func New(out io.Writer, prefix string, flag int) *Logger
type LogLevel int

const (
	Info LogLevel = iota
	Warning
	Error
)

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

type LogExtended struct {
	*log.Logger
	logLevel LogLevel // LogLevel это enum
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
