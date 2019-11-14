package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

func init() {
	if os.Getenv("ENV") == "DEV" {
		godotenv.Load("configs/Dev.yaml")
		initAppConfig()
		initNewRelicDevMode()
	} else {
		initAppConfig()
		initStackImpact()
		initLogentries()
		initRaven()
		initNewRelicProdMode()
	}
	initBusinessHours()
	initManualConfig()
	initDB()
	initMongoDB()
}

// ReadConfig reads config env variable into given struct
func ReadConfig(prefix string, i interface{}) error {
	err := envconfig.Process(prefix, i)
	return err
}
