package bootstrap

import (
	"github.com/jonboulle/clockwork"
	"github.com/sirupsen/logrus"
	"os"
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
	c.Logger = logger()

	return c
}

func logger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	return logger
}
