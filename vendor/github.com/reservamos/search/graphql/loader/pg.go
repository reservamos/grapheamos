package loader

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/graph-gophers/dataloader"
	"github.com/jinzhu/inflection"
	"github.com/reservamos/search/graphql/utils"
)

// DataLoaderQueryInts dataloader query for int keys
func DataLoaderQueryInts(keys dataloader.Keys, model interface{}, keyName string) string {
	return dlQuery(
		dlModelTableName(model),
		dlModelFields(model),
		dlSortTable(keys, func(k string) string {
			return k
		}),
		keyName,
	)
}

// DataLoaderQueryStrings dataloader query for string keys
func DataLoaderQueryStrings(keys dataloader.Keys, model interface{}, keyName string) string {
	return dlQuery(
		dlModelTableName(model),
		dlModelFields(model),
		dlSortTable(keys, func(k string) string {
			return fmt.Sprintf(`'%s'`, k)
		}),
		keyName,
	)
}

// DataLoaderOpenQuery dataloader query with option to open query
func DataLoaderOpenQuery(keys dataloader.Keys, fields string, table string, keyName string, keyType string) string {
	return dlQuery(
		fields,
		table,
		dlSortTable(keys, func(k string) string {
			if keyType == "string" {
				return fmt.Sprintf(`'%s'`, k)
			}
			return k
		}),
		keyName,
	)
}

func dlQuery(table string, fields string, sortTable string, fieldName string) string {
	return fmt.Sprintf(`
		SELECT %s
		FROM %s
			RIGHT OUTER JOIN %s
				ON sortTable.ref_id = %s
		ORDER BY sortTable.sort_order;
	`, fields, table, sortTable, fieldName)
}

func dlModelTableName(model interface{}) string {
	return utils.ToSnake(inflection.Plural(reflect.TypeOf(model).Name()))
}

func dlModelFields(model interface{}) string {
	rt := reflect.ValueOf(model).Type()
	fields := []string{}
	for i := 0; i < rt.NumField(); i++ {
		v, ok := rt.Field(i).Tag.Lookup("dl")
		if ok {
			fields = append(fields, v)
		}
	}
	return strings.Join(fields, ",")
}

func dlSortTable(keys dataloader.Keys, f func(string) string) string {
	values := make([]string, len(keys))
	for index, key := range keys {
		values[index] = fmt.Sprintf("(%s, %d)", f(key.String()), index)
	}
	return fmt.Sprintf(`(VALUES %s) as sortTable(ref_id, sort_order)`, strings.Join(values, ","))
}
