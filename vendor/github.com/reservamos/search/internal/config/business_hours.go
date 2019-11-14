package config

import (
	"time"

	"github.com/reservamos/search/perform/utils"
)

var bizConfig businessHours

type businessHours struct {
	WeekdayStart  int    `default:"7"`
	WeekdayFinish int    `default:"22"`
	WeekendStart  int    `default:"9"`
	WeekendFinish int    `default:"18"`
	Timezone      string `default:"America/Mexico_City"`
	Weekdays      []int  `default:"0,1,2,3,4,5,6"`
}

func initBusinessHours() {
	bizConfig = newBusinessHoursConfig()
}

func newBusinessHoursConfig() businessHours {
	var bh businessHours
	err := ReadConfig("BHOURS", &bh)
	if err != nil {
		panic(err)
	}
	return bh
}

// BusinessDurationUntil amount of business hours until limit - FOR DEBUG
func BusinessDurationUntil(limit time.Time) time.Duration {
	now := time.Now().In(loc())
	limit = limit.In(loc())
	businessDuration := 0 * time.Minute
	for date := now; date.Before(limit); date = date.AddDate(0, 0, 1) {
		businessDuration += dailyBusinessDuration(now, date, limit)
	}
	return businessDuration
}

// HasMinBusinessHoursUntil true if it has the minimum requirement of business hours. False if not
func HasMinBusinessHoursUntil(min int, limit time.Time) bool {
	now := time.Now().In(loc())
	limit = limit.In(loc())
	businessDuration := 0 * time.Minute
	minBusinessDuration := time.Duration(min) * time.Hour

	// Accelerate operation for longer periods of time
	if limit.Sub(now).Hours() > 96 {
		return true
	}
	for date := now; date.Before(limit); date = date.AddDate(0, 0, 1) {
		businessDuration += dailyBusinessDuration(now, date, limit)
		if businessDuration > minBusinessDuration {
			return true
		}
	}
	return false
}

func dailyBusinessDuration(now time.Time, date time.Time, limit time.Time) time.Duration {
	start := dateStartTime(date)
	finish := dateFinishTime(date)
	if utils.I.Include(bizConfig.Weekdays, int(date.Weekday())) {
		return (lesser(limit, finish).Sub(greater(start, now)))
	}
	return 0 * time.Minute
}

func greater(start time.Time, now time.Time) time.Time {
	if start.After(now) {
		return start
	}
	return now
}

func lesser(limit time.Time, finish time.Time) time.Time {
	if limit.After(finish) {
		return finish
	}
	return limit
}

func dateStartTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), dailyStartHour(d), 0, 0, 0, loc())
}

func dateFinishTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), dailyFinishHour(d), 0, 0, 0, loc())
}

func dailyStartHour(date time.Time) int {
	if isWeekday(date) {
		return bizConfig.WeekdayStart
	}
	return bizConfig.WeekendStart
}

func dailyFinishHour(date time.Time) int {
	if isWeekday(date) {
		return bizConfig.WeekdayFinish
	}
	return bizConfig.WeekendFinish
}

func isWeekday(date time.Time) bool {
	return date.Weekday() != time.Saturday && date.Weekday() != time.Sunday
}

func loc() *time.Location {
	loc, _ := time.LoadLocation(bizConfig.Timezone)
	return loc
}
