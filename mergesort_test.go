package main

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

// Test_parallel_mergesort uses goroutine to implement parallelly sorting
func Test_parallel_mergesort(t *testing.T) {

	input := rand.Perm(10000)

	convey.Convey("parallel merge sorting", t, func() {
		out := make(chan []int)
		go SortParallel(input, out)
		var output []int = <-out
		var err error = nil
		for i := 0; i < len(output)-1; i++ {
			if output[i] > output[i+1] {
				err = fmt.Errorf("worng order of output array: '%d' is not smaller than '%d'", output[i], output[i+1])
				break
			}
		}

		convey.So(err, convey.ShouldBeEmpty)

	})
}

func Test_mergesort(t *testing.T) {

	input := rand.Perm(10000)
	convey.Convey("non-parallelly merge sorting", t, func() {
		output := Sort(input)
		var err error = nil
		for i := 0; i < len(output)-1; i++ {
			if output[i] > output[i+1] {
				err = fmt.Errorf("worng order of output array: '%d' is not smaller than '%d'", output[i], output[i+1])
				break
			}
		}

		convey.So(err, convey.ShouldBeEmpty)
	})

}
