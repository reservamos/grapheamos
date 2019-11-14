package utils

import (
	"fmt"
	"strconv"
	"time"
)

// Duration returns a time.Duration substraction of two times
func Duration(oTime time.Time, oLocation string, dTime time.Time, dLocation string) time.Duration {
	return GetTimeTz(ParseTimeAsLocal(dTime), dLocation).Sub(GetTimeTz(ParseTimeAsLocal(oTime), oLocation))
}

// DateToLocal Gets a date in whichever format, extracts year, month and date and returns a time.Time
// object as a America/Mexico_City tz date
func DateToLocal(t time.Time) time.Time {
	tz, _ := time.LoadLocation("America/Mexico_City")
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, tz)
}

// DateToString formats time to string for graphql response
func DateToString(t time.Time) string {
	return t.Format("02-01-2006")
}

// DateToSQLString formats time to string for sql query
func DateToSQLString(t time.Time) string {
	return t.Format("2006-01-02")
}

func ServicesTimeToTime(s string) time.Time {
	parsed, _ := time.Parse("2006-01-02T15:04:05.999 MST", s+" UTC")
	return GetTimeTz(parsed, "America/Mexico_City")
}

func UnixToTime(u int64) time.Time {
	i, err := strconv.ParseInt(fmt.Sprintf("%v", u/1000), 10, 64)
	if err != nil {
		panic(err)
	}
	return time.Unix(i, 0)
}

func ParseTimeAsUTC(t time.Time) time.Time {
	t = t.UTC()
	loc, _ := time.LoadLocation("America/Mexico_City")

	return time.Date(
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
		t.Nanosecond(),
		loc,
	)
}

func ParseTimeAsLocal(t time.Time) time.Time {
	loc, _ := time.LoadLocation("America/Mexico_City")
	return t.In(loc)
}

func GetTimeTz(t time.Time, location string) time.Time {
	if location == "" {
		return t
	} else {
		loc, _ := time.LoadLocation(location)
		res := time.Date(
			t.Year(),
			t.Month(),
			t.Day(),
			t.Hour(),
			t.Minute(),
			t.Second(),
			t.Nanosecond(),
			loc,
		)

		return res
	}
}
