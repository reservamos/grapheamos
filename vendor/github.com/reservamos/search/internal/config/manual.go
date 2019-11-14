package config

import (
	"os"
	"strconv"
	"strings"
	"time"
)

// Manual config
var Manual ManualConfig

// ManualConfig Manual Lines configuration
type ManualConfig struct {
	Transporters   []string
	FilterEnabled  bool `default:"true"`
	DefaultWeekday int  `default:"6"`
	DefaultWeekend int  `default:"6"`
}

// ManualHours given and transporter and a time returns amount of hours needed to process a purchase
func ManualHours(transporter string, now time.Time) int {
	if now.Weekday() < time.Saturday {
		return ManualWeekdayHours(transporter)
	}
	return ManualWeekendHours(transporter)
}

// ManualWeekdayHours returns business hours needed on weekdays for given transporter
func ManualWeekdayHours(transporterAbbr string) int {
	raw := os.Getenv("MANUAL_" + strings.ToUpper(transporterAbbr) + "_WEEKDAY")
	if val, err := strconv.Atoi(raw); err == nil {
		return val
	}
	return Manual.DefaultWeekday
}

// ManualWeekendHours returns business hours needed on weekends for given transporter
func ManualWeekendHours(transporterAbbr string) int {
	raw := os.Getenv("MANUAL_" + strings.ToUpper(transporterAbbr) + "_WEEKEND")
	if val, err := strconv.Atoi(raw); err == nil {
		return val
	}
	return Manual.DefaultWeekend
}

func initManualConfig() {
	Manual = newManualConfig()
}

func newManualConfig() ManualConfig {
	var mc ManualConfig
	err := ReadConfig("manual", &mc)
	if err != nil {
		panic(err)
	}
	return mc
}
