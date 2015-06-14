// task2 project logger.go
// Logger is a singleton wrapper of Logger from go-logging/logging package

package controllers

import (
	"os"

	"github.com/op/go-logging"
)

var instance *logging.Logger

func GetLogger() *logging.Logger {
	if instance == nil {
		instance = logging.MustGetLogger("example")
		logBackend := logging.NewLogBackend(os.Stderr, "", 0)
		logFormat := logging.MustStringFormatter("%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level} %{message}%{color:reset}")
		logging.SetBackend(logging.NewBackendFormatter(logBackend, logFormat))
	}
	return instance
}
