package utils

import (
	"encoding/base64"
	"reflect"
	"strconv"
	"strings"

	graphql "github.com/graph-gophers/graphql-go"
)

// ToGlobalID transforms an int or string id into a graphql id
func ToGlobalID(nodeType string, id interface{}) graphql.ID {
	if reflect.TypeOf(id).Kind() == reflect.Int {
		return graphql.ID(toBase64(intIDBase(nodeType, id.(int))))
	}
	return graphql.ID(toBase64(stringIDBase(nodeType, id.(string))))
}

// ToGlobalEzID transforms an int or string id into an easy to input id
func ToGlobalEzID(nodeType string, id interface{}) graphql.ID {
	if reflect.TypeOf(id).Kind() == reflect.Int {
		return graphql.ID(intIDBase(nodeType, id.(int)))
	}
	return graphql.ID(stringIDBase(nodeType, id.(string)))
}

// FromGlobalIntID transforms graphql id to integer
func FromGlobalIntID(id string) (int, string) {
	id, name := FromGlobalStringID(id)
	idIntVal, _ := strconv.Atoi(id)
	return idIntVal, name
}

// FromGlobalStringID transforms graphql id to string
func FromGlobalStringID(id string) (string, string) {
	var nameID []string
	raw, err := fromBase64(id)
	if err != nil {
		nameID = strings.Split(id, ":")
	} else {
		nameID = strings.Split(raw, ":")
	}
	return nameID[1], nameID[0]
}

func intIDBase(nodeType string, id int) string {
	return nodeType + ":" + strconv.Itoa(id)
}

func stringIDBase(nodeType string, id string) string {
	return nodeType + ":" + id
}

func toBase64(s string) string {
	data := []byte(s)
	return base64.StdEncoding.EncodeToString(data)
}

func fromBase64(s string) (string, error) {
	result, err := base64.StdEncoding.DecodeString(s)
	return string(result), err
}
