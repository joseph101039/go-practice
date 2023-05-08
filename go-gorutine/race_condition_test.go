package main

import (
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"testing"
	"time"
)

func Test_race(t *testing.T) {

	access := make(chan int)
	shared := 0
	wg := &sync.WaitGroup{}
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT)
	wg.Add(2)
	go child1(access, &shared, wg, sig)
	go child2(access, &shared, wg, sig)

	time.Sleep(30 * time.Second)

	sig <- syscall.SIGINT
	sig <- syscall.SIGINT
	wg.Wait()
	close(access) // should not closed while sending
}

func child1(access chan int, shared *int, wg *sync.WaitGroup, sig <-chan os.Signal) {
	defer wg.Done()
	defer logError(nil)

	for {
		randSleep(5)
		writeShared(shared, 1)
		access <- 1                 // give access to child 2
		if _, ok := <-access; !ok { // blocking
			break
		}

		select {
		case s := <-sig:
			log.Printf("child %d recieve %s", 1, s)
			return
		default:
		}
	}
}

func child2(access chan int, shared *int, wg *sync.WaitGroup, sig <-chan os.Signal) {
	defer wg.Done()
	defer logError(func(e error) bool {
		return strings.HasPrefix(e.Error(), "send on closed channel")
	})

	for {
		if _, ok := <-access; !ok { // blocking
			break
		}

		randSleep(5)
		writeShared(shared, 2)
		access <- 2 // give access to child 2

		select {
		case s := <-sig:
			log.Printf("child %d recieve %s", 2, s)
			return
		default:
		}
	}
}

func logError(recoverCondition func(e error) bool) {
	if e := recover(); e != nil {
		log.Println(e)

		if err, ok := e.(error); ok && recoverCondition != nil {
			if !(recoverCondition)(err) {
				panic(err)
			}
		} else {
			panic(e)
		}
	}
}

func writeShared(shared *int, id int) {
	log.Printf("child %d write shared = %d", id, id)
	*shared = id
}

func randSleep(n int) {
	time.Sleep(time.Second * time.Duration(rand.Intn(n-1)+1))
}

func isChannelClosed(ch chan any) bool {
	select {
	case _, ok := <-ch:
		return !ok
	default:
	}
	return false
}
