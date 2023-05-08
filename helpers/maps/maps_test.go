package maps

import (
	"fmt"
	"testing"
)

func Test_FilterKeys(t *testing.T) {

	m := map[string]interface{}{
		"a": 1,
		"b": "2",
		"c": struct{ a int }{1},
	}

	allows := [...]string{"c", "b"} // array, not a slice

	fmt.Println(
		FilterKeys(m, allows[:]), // 要將固定長度陣列轉換為 slice，可以使用切片運算符 [:]
	)

}
