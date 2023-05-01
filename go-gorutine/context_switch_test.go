package main

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

func TestContextSWitch(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("err := %#v", err)
		}
	}()

	const loopCount = 1000
	var wg sync.WaitGroup
	wg.Add(loopCount)

	for i := 0; i < loopCount; i++ {
		wg.Add(1)
		go runLoopForever(&wg, i)
		// go runLoopWithSyscall(&wg, i)
	}
	wg.Wait()

}

func runLoopWithSyscall(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Printf("id: %d\n", id)
	for {
		// time.Sleep(100 * time.Millisecond)
		time.Sleep(1 * time.Second)
	}

}

func runLoopForever(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Printf("id: %d\n", id)
	for {
	}
}
