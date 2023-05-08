package slices

import (
	"golang.org/x/exp/constraints"
)

func ArrayMap[IN any, OUT any](callback func(IN) OUT, array []IN) []OUT {
	o := make([]OUT, len(array))
	for k, v := range array {
		o[k] = callback(v)
	}
	return o
}

func ArrayReduce[IN any, OUT any](array []IN, callback func(carry OUT, item IN) OUT, initial OUT) OUT {
	carry := initial
	for _, v := range array {
		carry = callback(carry, v)
	}

	return carry
}

// InSlice 檢查 needle 是否存在於 haystack 中 (模仿 php in_array())
func InArray[T comparable](needle T, haystack []T) bool {

	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false

}

// Max returns the maximum value in a slice
func Max[T constraints.Ordered](s []T) (m T) {
	m = s[0]
	for _, v := range s {
		if m < v {
			m = v
		}
	}
	return
}

func Min[T constraints.Ordered](s []T) (m T) {
	m = s[0]
	for _, v := range s {
		if m > v {
			m = v
		}
	}
	return
}

// ToUint 將 int slice 轉換成 uint slice
func ToUint(s []int) []uint {
	var result []uint
	for _, ele := range s {
		result = append(result, uint(ele))
	}
	return result
}
