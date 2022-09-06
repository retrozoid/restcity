package trac

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

func locatorString(structure interface{}) string {
	var (
		value      []string
		structType = reflect.TypeOf(structure)
		structVal  = reflect.ValueOf(structure)
	)
	for ; structVal.Kind() == reflect.Ptr; structVal = structVal.Elem() {
	}
	if structVal.Kind() == reflect.Struct {
		for i := 0; i < structVal.NumField(); i++ {
			name := structType.Field(i).Tag.Get("json")
			name = strings.TrimSuffix(strings.TrimSpace(name), ",omitempty")
			if name == "" || name == "-" {
				continue
			}
			val := structVal.Field(i)
			if !val.IsZero() {
				for ; val.Kind() == reflect.Ptr; val = val.Elem() {
				}
				any := val.Interface()
				if t, isTime := any.(time.Time); isTime {
					any = t.Format("20060102T150405-0700")
				}
				value = append(value, fmt.Sprintf("%s:%v", name, any))
			}
		}
		return strings.Join(value, ",")
	}
	return ""
}
