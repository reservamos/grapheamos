package config

// AppConfig app shared config variables
type AppConfig struct {
	ImageRepo              string
	NoCopyright            []string
	IntegrationsURL        string `required:"true"`
	IntegrationsAuthHeader string `required:"true"`
	TTLDefault             int    `default:"60"`
	InstallmentsMin        int    `default:"0"`
	Timezone               string `default:"America/Mexico_City"`
	Currency               string `default:"MXN"`
	Rome2RioKey            string
	GoogleMapsKey          string
}

// App accessor for global app config
var App = AppConfig{}

func initAppConfig() {
	err := ReadConfig("APP", &App)
	if err != nil {
		panic(err)
	}
}
