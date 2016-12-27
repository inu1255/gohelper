package gohelper

import (
	"log"
	"os"
	"strings"
	"time"
)

func NewLogger() *log.Logger {
	filename := os.Args[0]
	var logFile *os.File
	if len(os.Args) > 1 && os.Args[1] == "stdio" || strings.HasPrefix(filename, os.TempDir()) {
		logFile = os.Stdout
	} else {
		filename += ".log"
		os.Mkdir(filename, 0755)
		filename = strings.Join([]string{filename, "/", time.Now().Format("2006-01-02"), ".log"}, "")
		var err error
		logFile, err = os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			logFile = os.Stdout
		}
	}
	return log.New(logFile, "", log.Ltime|log.Lshortfile)
}
