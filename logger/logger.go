package logger

import (
	"os"

	logging "github.com/op/go-logging"
)

const (
	LOG_MODULE = "logging"

	DEBUG LogLevel = LogLevel(logging.DEBUG)
	INFO  LogLevel = LogLevel(logging.INFO)
)

type LogLevel logging.Level

var Logger *logging.Logger

func init() {
	Logger = logging.MustGetLogger(LOG_MODULE)
	format := logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	)
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	backendLeveled := logging.AddModuleLevel(backendFormatter)

	// Set the backends to be used.
	logging.SetBackend(backendLeveled)
}

func SetLevel(level LogLevel) {
	logging.SetLevel(logging.Level(level), LOG_MODULE)
}

func Debug(args ...interface{}) {
	Logger.Debug(args...)
}

func Debugf(str string, args ...interface{}) {
	Logger.Debugf(str, args...)
}

func Info(args ...interface{}) {
	Logger.Info(args...)
}

func Infof(str string, args ...interface{}) {
	Logger.Infof(str, args...)
}
