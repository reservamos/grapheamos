package utils

import (
	"time"
)

var spanishMonths map[time.Month]string

func init() {
	spanishMonths = map[time.Month]string{
		time.January:   "Ene",
		time.February:  "Feb",
		time.March:     "Mar",
		time.April:     "Abr",
		time.May:       "May",
		time.June:      "Jun",
		time.July:      "Jul",
		time.August:    "Ago",
		time.September: "Sep",
		time.October:   "Oct",
		time.November:  "Nov",
		time.December:  "Dic",
	}
}

func TimeLongSlug(t time.Time) string {
	return t.Format("2") + spanishMonths[t.Month()] + t.Format("061504")
}

func DateLongSlug(t time.Time) string {
	return t.Format("02-") + spanishMonths[t.Month()] + t.Format("-2006")
}

func ApiTime(t time.Time) string {
	return t.Format("02-01-2006")
}
