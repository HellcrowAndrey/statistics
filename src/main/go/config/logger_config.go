package config

import (
	"log"
	"os"
)

var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

type Logger struct {
	config *Config
}

func NewLogger(config *Config) *Logger {
	return &Logger{config: config}
}

func (logger *Logger) InitLogger() {
	if logger.config.IsLoggerFile {
		file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			panic(err.Error())
		}
		Info = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		Warning = log.New(file, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
		Error = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	} else {
		Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		Warning = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
		Error = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	}
}
