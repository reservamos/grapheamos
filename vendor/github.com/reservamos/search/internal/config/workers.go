package config

import "fmt"

// WorkerConfig configuration for worker app
type WorkerConfig struct {
	Concurrency int      `required:"true"`
	Queues      []string `required:"true"`
	Discover    int      `required:"true"`
}

// NewWorkerConfig WorkerConfig constructor
func NewWorkerConfig() *WorkerConfig {
	var c WorkerConfig
	err := ReadConfig("WORKER", &c)
	if err != nil {
		panic(err)
	}
	if c.Discover == 0 {
		panic(fmt.Errorf("Discover duration not set"))
	}
	return &c
}
