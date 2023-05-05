package maths

import "golang.org/x/exp/constraints"

func Max[T constraints.Ordered](v ...T) (o T) {
	o = v[0]
	for _, element := range v {
		if element > o {
			o = element
		}
	}
	return
}

func Min[T constraints.Ordered](v ...T) (o T) {
	o = v[0]
	for _, element := range v {
		if element < o {
			o = element
		}
	}
	return
}
