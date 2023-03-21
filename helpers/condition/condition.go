package condition

import "reflect"

func If(condition bool, valueIfTrue string, valueIfFalse string) string {
	if condition {
		return valueIfTrue
	} else {
		return valueIfFalse
	}
}

func IfNull(value interface{}, altValue interface{}) interface{} {
	if reflect.ValueOf(value).IsNil() {
		return value
	} else {
		return altValue
	}
}
