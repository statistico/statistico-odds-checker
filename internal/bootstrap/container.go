package bootstrap

import (
	"github.com/evalphobia/logrus_sentry"
	"github.com/jonboulle/clockwork"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

type Container struct {
	Config *Config
	Clock  clockwork.Clock
	Logger *logrus.Logger
}

func BuildContainer(config *Config) Container {
	c := Container{
		Config: config,
	}

	c.Clock = clockwork.NewRealClock()
	c.Logger = logger(config)

	return c
}

func logger(config *Config) *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)

	tags := map[string]string{
		"application": "statistico-odds-checker",
	}

	levels := []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
	}

	hook, err := logrus_sentry.NewWithTagsSentryHook(config.Sentry.DSN, tags, levels)

	if err == nil {
		hook.Timeout = 20 * time.Second
		hook.StacktraceConfiguration.Enable = true
		hook.StacktraceConfiguration.IncludeErrorBreadcrumb = true
		logger.AddHook(hook)
	}

	return logger
}
