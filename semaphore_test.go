package main

import "testing"


type Request struct {
    args        []int
    f           func([]int) int
    resultChan  chan int
}



func Test_semaphore(t *testing.T) {

	const maxRequest int = 20
	var sem = make(chan *Request, maxRequest)


	

}

func requestHandler(r *Request, sem chan *Request)  {
	sem <
}