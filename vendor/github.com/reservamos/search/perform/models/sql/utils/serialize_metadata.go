package utils

import (
	"encoding/json"
	"reflect"
	"regexp"

	"github.com/jinzhu/gorm/dialects/postgres"
)

// SerializeMetadata receives a string map of whatever and serializes for postgres HStore
func SerializeMetadata(md map[string]interface{}) postgres.Hstore {
	meta := postgres.Hstore{}
	for k, v := range md {
		if reflect.TypeOf(v).Kind() == reflect.String {
			val := v.(string)
			meta[k] = &val
		} else {
			if js, err := json.Marshal(v); err == nil {
				re := regexp.MustCompile("\":")
				jst := re.ReplaceAllString(string(js), "\" => ")
				meta[k] = &jst
			}
		}
	}
	return meta
}
