package main

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

/**
goroutine 3697 [semacquire]:
internal/poll.runtime_Semacquire(0x10?)
        /usr/local/go/src/runtime/sema.go:67 +0x27
internal/poll.(*fdMutex).rwlock(0xc000022120, 0xd0?)
        /usr/local/go/src/internal/poll/fd_mutex.go:154 +0xd2
internal/poll.(*FD).writeLock(...)
        /usr/local/go/src/internal/poll/fd_mutex.go:239
internal/poll.(*FD).Write(0xc000022120, {0xc000fc33d0, 0x9, 0x10})
        /usr/local/go/src/internal/poll/fd_unix.go:370 +0x72
os.(*File).write(...)
        /usr/local/go/src/os/file_posix.go:48
os.(*File).Write(0xc000012018, {0xc000fc33d0?, 0x9, 0xc000fbc7a8?})
        /usr/local/go/src/os/file.go:175 +0x65
fmt.Fprintf({0x1517f20, 0xc000012018}, {0x147502d, 0x7}, {0xc000fbc7a8, 0x1, 0x1})
        /usr/local/go/src/fmt/print.go:225 +0x9b
fmt.Printf(...)
        /usr/local/go/src/fmt/print.go:233
goroutine.runLoopForever(0x0?)
        /Applications/XAMPP/xamppfiles/htdocs/go-gorotine/context_switch_test.go:20 +0x5d
created by goroutine.TestContextSWitch
        /Applications/XAMPP/xamppfiles/htdocs/go-gorotine/context_switch_test.go:14 +0x65


*/

func TestContextSWitch(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("err := %#v", err)
		}
	}()

	const loopCount = 5
	var wg sync.WaitGroup
	//wg.Add(loopCount)
	for i := 0; i < loopCount; i++ {
		wg.Add(1)
		// go runLoopForever(&wg, i)
		go runLoopWithSyscall(&wg, i)
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
