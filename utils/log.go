package utils

import (
	"log"
	"os"
)

func Setlog(logfilename string, flag int, logger *log.Logger) *log.Logger {
	logFile, err := os.OpenFile(logfilename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Println(err)
		return nil
	}
	logger = log.New(logFile, "", flag)
	logger.SetOutput(logFile)
	logger.SetFlags(flag)
	return logger
}
