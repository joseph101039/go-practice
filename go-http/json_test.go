package gohttp_test

import (
	"fmt"
	"goroutine/helpers/maps"
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	Field   string `json:"field"`
	Another int
}

func Test_json(t *testing.T) {
	v, _ := maps.ToJson(1)
	fmt.Println(v)
	fmt.Println(maps.ToJson(1.11))
	fmt.Println(maps.ToJson("abcde"))
	fmt.Println(maps.ToJson("{\"a\": 1}"))
	fmt.Println(maps.ToJson(map[string]interface{}{
		"a": "b",
		"c": 1.11,
	}))

	fmt.Println(maps.ToJson(test{Field: "A", Another: 1}))
	assert.IsType(t, "", v) // v 是字串
	assert.IsType(t, 1, v)  // v 是整數

}
