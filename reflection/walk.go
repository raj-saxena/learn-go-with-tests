package reflection

import (
	"reflect"
)

func getValue(i interface{}) (val reflect.Value) {
	val = reflect.ValueOf(i)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return
}

func Walk(input interface{}, f func(string)) {
	val := getValue(input)

	if val.Kind() == reflect.Slice || val.Kind() == reflect.Array {
		for i := 0; i < val.Len(); i++ {
			Walk(val.Index(i).Interface(), f)
		}
		return
	}

	for i := 0; i < val.NumField(); i++ {
		curr := val.Field(i)
		switch curr.Kind() {
		case reflect.String:
			f(curr.String())
		case reflect.Struct, reflect.Ptr:
			Walk(curr.Interface(), f)
		}
	}
}
