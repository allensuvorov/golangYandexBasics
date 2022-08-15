package main

import "log"

// func New(out io.Writer, prefix string, flag int) *Logger

type LogExtended struct {
	*log.Logger
	logLevel LogLevel // LogLevel это enum
}

func main() {
	logger := NewLogExtended()
	logger.SetLogLevel(LogLevelWarning)
	logger.Infoln("Не должно напечататься")
	logger.Warnln("Hello")
	logger.Errorln("World")
	logger.Println("Debug")
}
