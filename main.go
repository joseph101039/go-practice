package main

/*
#include <stdio.h>
#include <stdlib.h>
void printint(int v) {
	printf("print: %d\n", v);
}

void printstring(char* v) {
	printf("%s\n", v);
}
*/
import "C" // cgo 需要單獨 import, C 程式寫在註解中, 與 import C 之間不能有任何換行間隔
import (
	"flag"
	"goroutine/helpers/goerror"
	"log"
	"math/rand"
	"strconv"
	"unsafe"
)

func main() {

	// error handling
	defer func() {
		if err := recover(); err != nil {
			goerror.GetStackTrace(err)
		}
	}()

	cgo_cal(5)
	// return

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

// cgo_call calls the user-defined C functions, 如果 golang 缺少的 function 可以呼叫 C 但是需要考慮跨平台和相容性等問題
func cgo_cal(v int) {
	a1 := C.CString("Hello world")
	C.printstring(a1)
	C.free(unsafe.Pointer(a1)) // char pointer. the C string is allocated in the C heap using malloc.

	C.printint(C.int(v))

}
