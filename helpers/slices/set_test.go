package slices

import (
	"fmt"
	"math/rand"
	"testing"
)

func Test_sort(t *testing.T) {

	a := rand.Perm(100)
	fmt.Printf("%#v\n\n", a)
	Sort(a)

	fmt.Printf("%#v\n", a)

}
