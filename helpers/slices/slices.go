package slices

import (
	"errors"
	"reflect"
)

// InSlice 檢查 needle 是否存在於 haystack 中 (模仿 php in_array())
// reference: https://www.golangprograms.com/how-to-check-if-an-item-exists-in-slice-in-golang.html
func InSlice(needle interface{}, haystack interface{}) (bool, error) {
	h := reflect.ValueOf(haystack)
	if h.Kind() != reflect.Slice {
		return false, errors.New("haystack is not a slice")
	}

	for i := 0; i < h.Len(); i++ {
		if h.Index(i).Interface() == needle {
			return true, nil
		}

	}
	return false, nil
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
