package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/reservamos/search/perform/utils"
)

// StringsForWhere given a slice of strings ["a", "b"] returns "'a','b'"
func StringsForWhere(sts []string) string {
	return strings.Join(utils.Map(sts, func(s string) string {
		return fmt.Sprintf("'%s'", s)
	}), ",")
}

func idsToString(ids []int) string {
	result := make([]string, len(ids))
	for i, id := range ids {
		result[i] = strconv.Itoa(id)
	}
	return strings.Join(result, ",")
}
