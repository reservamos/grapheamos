package config

import (
	raven "github.com/getsentry/raven-go"
)

type sentryConfig struct {
	Dsn string `required:"true"`
	Env string `envconfig:"env"`
}

func newSentryConfig() *sentryConfig {
	var c sentryConfig
	err := ReadConfig("SENTRY", &c)
	if err != nil {
		panic(err)
	}
	return &c
}

func initRaven() {
	c := newSentryConfig()
	raven.SetDSN(c.Dsn)
	raven.SetEnvironment(c.Env)
}
