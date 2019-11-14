package utils

import (
	"time"

	"github.com/reservamos/search/internal/config"
)

// DateToLocal Gets a date in whichever format, extracts year, month and date and returns a time.Time
// object as a America/Mexico_City tz date
func DateToLocal(t time.Time) time.Time {
	tz, _ := time.LoadLocation(config.App.Timezone)
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, tz)
}

// TimeToLocal returns time in Mexico City Timezone
func TimeToLocal(t time.Time) time.Time {
	loc, _ := time.LoadLocation(config.App.Timezone)
	return t.In(loc)
}

// DateToString formats time to string for graphql response
func DateToString(t time.Time) string {
	return t.Format("02-01-2006")
}

// TimeToString formats time to string for graphql response
func TimeToString(t time.Time) string {
	return t.Format("02-01-2006 15:04:05 -0700")
}

// NullableTimeToLocalString returns local time in string when possible, nil when not
func NullableTimeToLocalString(t time.Time) *string {
	if !t.IsZero() {
		result := TimeToString(TimeToLocal(t))
		return &result
	}
	return nil
}

// NullableTimeToString returns string when possible, nil when not
func NullableTimeToString(t time.Time) *string {
	if !t.IsZero() {
		result := TimeToString(t)
		return &result
	}
	return nil
}
