package trac

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

var timeKind = reflect.TypeOf(time.Time{})

const omitempty = ",omitempty"

func toString(val reflect.Value) string {
	switch val.Kind() {
	case reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%d", val.Int())
	case reflect.Bool:
		return fmt.Sprintf("%t", val.Bool())
	case reflect.String:
		return val.String()
	case reflect.Struct:
		if val.Type() == timeKind {
			return val.Interface().(time.Time).Format("20060102T150405-0700")
		} else {
			return fmt.Sprintf("(%s)", EncodeLocator(val.Interface()))
		}
	default:
		return ""
	}
}

func EncodeLocator(structure interface{}) string {
	var (
		value      []string
		structType = reflect.TypeOf(structure)
		structVal  = reflect.ValueOf(structure)
	)
	for ; structVal.Kind() == reflect.Ptr; structVal = structVal.Elem() {
	}
	if structVal.Kind() == reflect.Struct {
		var raw = structVal.FieldByName("Raw")
		if raw.IsValid() && !raw.IsZero() {
			return raw.String()
		}
		for i := 0; i < structVal.NumField(); i++ {
			name := strings.TrimSpace(structType.Field(i).Tag.Get("json"))
			isOmitempty := strings.HasSuffix(name, omitempty)
			if isOmitempty {
				name = strings.TrimSuffix(name, omitempty)
			}
			if name == "" || name == "-" {
				continue
			}
			val := structVal.Field(i)
			if !val.IsZero() || !isOmitempty {
				for ; val.Kind() == reflect.Ptr; val = val.Elem() {
				}
				if val.Kind() == reflect.Slice || val.Kind() == reflect.Array {
					// item:(<locator1>),item:(<locator2>...)
					var slice []string
					for n := 0; n < val.Len(); n++ {
						slice = append(slice, fmt.Sprintf("%s:(%s)", name, toString(val.Index(n))))
					}
					value = append(value, strings.Join(slice, ","))
				} else {
					value = append(value, fmt.Sprintf("%s:%s", name, toString(val)))
				}
			}
		}
		return strings.Join(value, ",")
	}
	return ""
}
