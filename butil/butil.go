package butil

import (
	"fmt"
	"reflect"
)

// PrintReadableTypeValue returns a string representation of the type and value of the given value.
// Todo: to compatible with struct type that contains pointer fields.
func PrintReadableTypeValue(value interface{}) string {
	v := reflect.ValueOf(value)
	t := v.Type()

	if v.Kind() != reflect.Ptr {
		return fmt.Sprintf("(%s)(%v)", t, v.Interface())
	}

	var typeStr, valueStr string
	for elem := v; ; elem = elem.Elem() {
		if elem.Kind() != reflect.Ptr {
			typeStr += elem.Type().String()
			valueStr = fmt.Sprintf("(%+v)", elem.Interface())
			break
		} else if elem.IsNil() {
			typeStr += elem.Type().String()
			valueStr = "(nil)"
			break
		}
		typeStr += "*"
	}

	return fmt.Sprintf("(%s)%s", typeStr, valueStr)
}
