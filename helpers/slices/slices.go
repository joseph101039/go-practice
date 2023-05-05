package slices

import (
	"errors"
	"reflect"

	"golang.org/x/exp/constraints"
)

// InSlice 檢查 needle 是否存在於 haystack 中 (模仿 php in_array())
// reference: https://www.golangprograms.com/how-to-check-if-an-item-exists-in-slice-in-golang.html
func InSlice[T comparable](needle T, haystack []T) bool {

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

// Reverse 反轉元素位置
func Reverse(s interface{}) {
	h := reflect.ValueOf(s)
	if h.Kind() != reflect.Slice {
		panic(errors.New("not a slice"))
	}

	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

// ToUint 將 int slice 轉換成 uint slice
func ToUint(s []int) []uint {
	var result []uint
	for _, ele := range s {
		result = append(result, uint(ele))
	}
	return result
}
