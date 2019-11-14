package config

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bsphere/le_go"
)

// Log logger global
var Log *le_go.Logger

// ServerID global
var ServerID int

type logenetriesConfig struct {
	Token string `required:"true"`
}

func newLogenetriesConfig() *logenetriesConfig {
	var c logenetriesConfig
	err := ReadConfig("LOGENTRIES", &c)
	if err != nil {
		panic(err)
	}
	return &c
}

func initLogentries() {
	c := newLogenetriesConfig()

	le, err := le_go.Connect(c.Token)
	if err != nil {
		fmt.Println(err)
	}

	generateServerID()
	Log = le
}

// Generate new seed to distinguish each server
func generateServerID() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	ServerID = r1.Intn(100)
}
