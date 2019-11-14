package utils

import (
	"sort"
	"time"
)

// PromoCode interface
type PromoCode interface {
	Percent() int
	Validate(time.Time) bool
	CodeString() string
}

type byPercent []PromoCode

func (s byPercent) Len() int {
	return len(s)
}
func (s byPercent) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byPercent) Less(i, j int) bool {
	return s[i].Percent() < s[j].Percent()
}

// BestPromoCode returns the best promocode of a set related to a specific date
func BestPromoCode(codes []PromoCode, date time.Time) string {
	if len(codes) != 0 {
		validCodes := []PromoCode{}
		for _, code := range codes {
			if code.Validate(date) {
				validCodes = append(validCodes, code)
			}
		}
		if len(validCodes) > 0 {
			sort.Sort(byPercent(validCodes))
			return validCodes[len(validCodes)-1].CodeString()
		}
	}
	return ""
}
