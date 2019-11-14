package config

import (
	stackimpact "github.com/stackimpact/stackimpact-go"
)

type stackimpactConfig struct {
	Key  string `require:"true"`
	Name string `required:"true"`
	Env  string `envconfig:"env"`
}

func newStackimpactConfig() *stackimpactConfig {
	var c stackimpactConfig
	err := ReadConfig("STACKIMPACT", &c)
	if err != nil {
		panic(err)
	}
	return &c
}

func initStackImpact() {
	c := newStackimpactConfig()
	stackimpact.Start(stackimpact.Options{
		AgentKey:       c.Key,
		AppName:        c.Name,
		AppEnvironment: c.Env,
	})
}
