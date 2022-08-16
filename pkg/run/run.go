package run

import (
	"github.com/sirupsen/logrus"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Args struct {
	ServiceName string
	ConfigFile  string
	MetricsAddr string
	LogLevel    logrus.Level
}

func (a Args) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.ConfigFile, validation.Required),
		validation.Field(&a.MetricsAddr, validation.Required),
	)
}

func LogLevel(level string) logrus.Level {
	switch level {
	case "DEBUG":
		return logrus.DebugLevel
	case "INFO":
		return logrus.InfoLevel
	case "WARN":
		return logrus.WarnLevel
	case "ERROR":
		return logrus.ErrorLevel
	}

	logrus.Errorf(`unknown log level %q, using "ERROR" by default`, level)

	return logrus.ErrorLevel
}
