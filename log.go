package main

import (
	"io/ioutil"
	"os"

	"github.com/Sirupsen/logrus"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

const (
	logFile    = "/tmp/logfile.log"
	maxAge     = 5   // days
	maxSize    = 500 // MB
	maxBackups = 20
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(ioutil.Discard) // no stderr
	file := os.Getenv("LOG_FILE")
	if file == "" {
		file = logFile
	}
	logrus.AddHook(NewLJHook(file, maxAge, maxSize, maxBackups))
}

type LJHook struct {
	w *lumberjack.Logger
}

func NewLJHook(file string, age, size, backups int) logrus.Hook {
	return LJHook{
		w: &lumberjack.Logger{
			Filename:   file,
			MaxAge:     age,
			MaxSize:    size,
			MaxBackups: backups,
			LocalTime:  true,
			Compress:   true,
		},
	}
}

func (_ LJHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (l LJHook) Fire(entry *logrus.Entry) (err error) {
	line, err := entry.String()
	if err != nil {
		return err
	}
	l.w.Write([]byte(line))
	return nil
}
