package slices

import (
	"errors"
	"reflect"

	"golang.org/x/exp/constraints"
)

// Unique returns a unique subset of the slice provided.
func Unique[T comparable](input []T) []T {
	u := []T{}
	m := make(map[T]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}
	return u
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

// Sort 針對排序
func Sort[T constraints.Ordered](s []T) {
	var tmp T
	for i := 1; i < len(s)-1; i++ {
		for j := 0; j < len(s)-i; j++ {
			if s[j] > s[j+1] {
				tmp = s[j]
				s[j] = s[j+1]
				s[j+1] = tmp
			}
		}
	}
	return
}

// Intersect 取所有 slice 交集
func Intersect[T comparable](slice []T, slices ...[]T) []T {

	hash := make(map[T]struct{}, len(slice))
	for _, v := range slice {
		hash[v] = struct{}{}
	}

	// remove the keys if not in other slices
	for i := 0; i < len(slices); i++ {
		for _, v := range slices[i] {
			if _, ok := hash[v]; !ok {
				delete(hash, v)
			}
		}
	}

	set := make([]T, len(hash))
	for v, _ := range hash {
		set = append(set, v)
	}

	return set
}

func Diff[T comparable](slice []T, slices ...[]T) []T {
	hash := make(map[T]struct{}, len(slice))
	for _, v := range slice {
		hash[v] = struct{}{}
	}

	// remove the keys which are in other slices
	for i := 0; i < len(slices); i++ {
		for _, v := range slices[i] {
			if _, ok := hash[v]; ok {
				delete(hash, v)
			}
		}
	}

	set := make([]T, len(hash))
	for v, _ := range hash {
		set = append(set, v)
	}

	return set
}
