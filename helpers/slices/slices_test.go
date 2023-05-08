package slices

import (
	"fmt"
	"testing"
)

func Test_arraymap(t *testing.T) {

	integers := []int{1, 2, 3, 4}

	r1 := ArrayMap(func(v int) string {
		return fmt.Sprint(v)
	}, integers)

	fmt.Println(r1)

	r2 := ArrayMap(stoi, integers)
	fmt.Println(r2)

}

func Test_arrayreduce(t *testing.T) {
	integers := []int{1, 2, 3, 4}

	r1 := ArrayReduce(integers, func(carry string, item int) string {
		return fmt.Sprintf("%s,%d", carry, item)
	}, "0")

	fmt.Println(r1)
}

func TestMax(t *testing.T) {
	fmt.Println(Max([]int{1, 2, 3}))
	v := []float32{1, 2, 3}
	fmt.Println(Max(v))
}

func stoi(v int) string {
	return fmt.Sprint(v)
}
