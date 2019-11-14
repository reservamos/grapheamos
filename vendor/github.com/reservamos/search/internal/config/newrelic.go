package config

import (
	"log"
	"os"

	newrelic "github.com/newrelic/go-agent"
)

var app newrelic.Application

//NewRelicConfig encapsulates new relic configuration
type NewRelicConfig struct {
	App newrelic.Application
}

func GetNewRelicInstance() newrelic.Application {
	return app
}

func initNewRelicDevMode() {
	initNewRelic(false)
}

func initNewRelicProdMode() {
	initNewRelic(true)
}

func initNewRelic(isEnabled bool) {
	var err error
	config := newrelic.NewConfig(getAppName(), getRelicKey())
	config.Enabled = isEnabled
	app, err = newrelic.NewApplication(config)
	if err != nil {
		log.Fatalln("NewRelic could not be linked")
	}
}

func getRelicKey() string {
	return os.Getenv("NEW_RELIC_LICENSE_KEY")
}

func getAppName() string {
	var appName string
	if os.Getenv("NEW_RELIC_APP_NAME") == "" {
		appName = "reservamos-search-w-no-name"
		log.Println("Using default app name")
	} else {
		appName = os.Getenv("NEW_RELIC_APP_NAME")
	}
	return appName
}
