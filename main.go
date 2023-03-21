package main

import (
	"flag"
	"goroutine/helpers/goerror"
	"log"
	"math/rand"
	"strconv"
)

func main() {

	// error handling
	defer func() {
		if err := recover(); err != nil {
			goerror.GetStackTrace(err)
		}
	}()

	var err error
	var size int = 100

	if flag.NArg() > 0 {
		size, err = strconv.Atoi(flag.Arg(0))
		goerror.Fatal(err)
	}

	input := rand.Perm(size)

	// 1
	output := Sort(input)

	// 2
	ch := make(chan []int)
	go SortParallel(input, ch)

	log.Printf("input: %#v\noutput: %#v\nparallel output: %#v", input, output, <-ch)
}
