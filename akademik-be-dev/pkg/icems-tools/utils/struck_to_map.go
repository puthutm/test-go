package utils

import (
	"fmt"
	"reflect"
)

func StructToMapString(input any) map[string]string {
	result := make(map[string]string)
	v := reflect.ValueOf(input)
	t := reflect.TypeOf(input)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
		t = t.Elem()
	}
	if v.Kind() != reflect.Struct {
		return result
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)
		key := fieldType.Tag.Get("form")
		if key == "" {
			key = fieldType.Name
		}
		if field.CanInterface() {
			result[key] = fmt.Sprintf("%v", field.Interface())
		}
	}
	return result
}
