package main

import (
	"fmt"
	"reflect"
	"testing"
)

/**
@link test command: https://gist.github.com/sebnyberg/0d39e5ae2e4d398f267a8c3383f49df3

1. normal test
go test -v -run=^Test_findPrimes$

2. benchmark a test with profiles and trace output (bypass all -run test)
go test -v -run=^$ -bench=findPrimesParallelNonWait  \
	-cpuprofile=profiles/cpu.out \
	-memprofile=profiles/mem.out \
	-trace=profiles/trace.out \
	-blockprofile=profiles/block.out

3. dump profile content
 go tool pprof -text goroutine.test.exe profiles/mem.out

4. dump another format require "brew install graphviz"
go tool pprof -svg goroutine.test.exe profiles/mem.out > profiles/mem.out.svg

5. 查看 trace
go tool trace trace.out
*/

const startNumber int = 60001
const endNumber int = 90000

// Test_findPrimes find some prime numbers non-parallelly
func Test_findPrimes(t *testing.T) {
	var isPrime bool
	for num := startNumber; num <= endNumber; num+=2 {
		if isPrime = isPrimeNumber(num); isPrime {
			fmt.Printf("%d ", num)
		}
	}

}

// Test_findPrimes uses goroutine to find some prime numbers
func Test_findPrimesParallel(t *testing.T) {
	findPrimesParallelTest(nil)
}

// Test_findPrimesParallelNonWait use select channels to reduc blocking time
func Test_findPrimesParallelNonWait(t *testing.T) {
	findPrimesParallelTestNonWait(nil)

}

func Test_findPrimesParallelChanMultiplex(t *testing.T) {
	// todo implement: https://go.dev/doc/effective_go#chan_of_chan
}

func Benchmark_findPrimesParallel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findPrimesParallelTest(b)
	}
}

func Benchmark_findPrimesParallelNonWait(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findPrimesParallelTestNonWait(b)
	}
}

func findPrimesParallelTestNonWait(b *testing.B) {

	// 初始化 channels 的地方不要測試
	var ch [int(endNumber - startNumber + 1)/2]chan bool // fixed length allocation use array instead of slice
	for i := 0; i < len(ch); i++ {
		ch[i] = make(chan bool)
	}

	if b != nil {
		b.StartTimer()
	}

	for num := startNumber; num <= endNumber; num++ {
		// increment the waitGroup counter
		go isPrimeNumberCh(num, ch[num-startNumber])
	}

	cases := make([]reflect.SelectCase, len(ch))
	for i, c := range ch {
		cases[i] = reflect.SelectCase{
			Dir:  reflect.SelectRecv, // direction
			Chan: reflect.ValueOf(c), // channel
		}
	}

	// primes := make(map[int]bool)
	for num := startNumber; num <= endNumber; num++ {
		// 讀取優先完成的 channels
		chosen, value, ok := reflect.Select(cases)
		if !ok {
			// The chosen channel has been closed, so zero out the channel to disable the case
			cases[chosen].Chan = reflect.ValueOf(nil)
			continue
		}

		if value.Bool() {
			// do nothing
		}

		// primes[chosen+startNumber] = value.Bool()
		// if value.Bool() { // is a prime
		// 	fmt.Printf("%d ", chosen+startNumber)
		// }
	}

	// for num, isPrime := range primes {
	// 	if isPrime {
	// 		fmt.Printf("%d ", num)
	// 	}
	// }

	if b != nil {
		b.StopTimer()
	}
}

func findPrimesParallelTest(b *testing.B) {

	// 初始化 channels 的地方不要測試
	var ch [endNumber - startNumber + 1]chan bool // fixed length allocation use array instead of slice
	for i := 0; i < len(ch); i++ {
		ch[i] = make(chan bool)
	}

	if b != nil {
		b.StartTimer()
	}

	for num := startNumber; num <= endNumber; num++ {
		// increment the waitGroup counter
		go isPrimeNumberCh(num, ch[num-startNumber])
	}

	for num := startNumber; num <= endNumber; num++ {
		if <-ch[num-startNumber] {
			fmt.Printf("%d ", num)
		}
	}

	if b != nil {
		b.StopTimer()
	}
}

// isPrimeNumberCh only accespts a channel for sending value
func isPrimeNumberCh(number int, ch chan<- bool) {
	ch <- isPrimeNumber(number)
}

func isPrimeNumber(number int) bool {
	if number%2 == 0 && number != 2 {
		return false
	} else {
		for i := 3; i < number; i += 2 {
			if number%i == 0 {
				return false
			}
		}

		return true
	}
}
