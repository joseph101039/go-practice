package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"testing"
	"time"
)

type Request struct {
	args       []interface{}
	f          func(...interface{}) interface{}
	resultChan chan interface{}
}


func Test_bufferedChannel(t *testing.T) {
	concurrency := runtime.GOMAXPROCS(0)
	bufferedChannel(t, concurrency, 100)
}


func Benchmark_bufferedChannel(b *testing.B) {
	bufferedChannel(b, 20, 100)
}

func parallel() {
	// cpuNum := runtime.NumCPU()
	


}

// semaphore 
// for  rate-limited, parallel, (non-blocking ?) RPC system, and there's not a mutex in sight.
func bufferedChannel(t testing.TB, concurrency int, testingLoop int) {

	t.Logf("%d concurrency handlers and total %d requests\n", concurrency, testingLoop)
	
	resultChan := NewResultChannels(concurrency)	
	reqChan := make(chan *Request, concurrency)

	DispatchHandlers(reqChan, concurrency)
	t.Logf("%d handlers are dispatched\n", concurrency)



	var cases = make([]reflect.SelectCase, concurrency)

	var process = func(parameters ...interface{}) (sum interface{}) {
		r := 0
		for _, e := range parameters {
			r += e.(int)
		}

		time.Sleep(time.Second * time.Duration(rand.Intn(3) + 1))  // sleep random seconds
		return r
	}

	// 模擬處理結果 (can be use as aggregate algorithm)
	var resultCallback = func(index int, result reflect.Value) {
		sum := result.Interface().(int)
		fmt.Printf("result %d is output: %d\n", index, sum)

	}


	var result reflect.Value
	for i := 0; i < testingLoop; i++ {	
		
		current := i % concurrency // default
		

		// get result if the request channel is full
		if len(reqChan) == cap(reqChan) {
			current, result, _ = FetchResult(&cases)
			resultCallback(current, result) 
		}



		EnqueueRequest(&(cases[current]), reqChan, &Request{
			args: []interface{} {
					rand.Intn(10),
					rand.Intn(10), 
					rand.Intn(10),
				},
			f:          process,
			resultChan: resultChan[current],
		})


		t.Logf("the %d th request are send", i + 1)
	}

	// todo some results may not be fetched

	// 取出剩餘結果
	for len(reqChan) != 0 {
		chosen, result, _ := FetchResult(&cases)
		resultCallback(chosen, result) 
	}

	
	close(reqChan)  // to make all recievers finishing waiting
	t.Logf("There are %d goroutines are currently running", runtime.NumGoroutine())
}


func EnqueueRequest(rcase *reflect.SelectCase, reqChan chan <- *Request, req *Request) {
	// push a new request case into queue
	*rcase = reflect.SelectCase{
		Dir:reflect.SelectRecv,
		Chan: reflect.ValueOf(req.resultChan),
	}

	reqChan <- req
}

func FetchResult(cases *[]reflect.SelectCase) (chosen int, result reflect.Value, err error) {
	chosen, result, ok  := reflect.Select(*cases)
			
	if !ok {
		// The chosen channel has been closed, so zero out the channel to disable the case
		(*cases)[chosen].Chan = reflect.ValueOf(nil)
		return -1, reflect.Value{}, fmt.Errorf("[Error] cannot receive the result of case %d", chosen)
	}

	// 模擬處理結果 (can be use as aggregate algorithm)
	return chosen, result, nil
}




func NewResultChannels(concurrency int) []chan interface{} {
	resultChan := make([]chan interface{}, concurrency)
	for i := range resultChan {
		resultChan[i] = make(chan interface{})
	}
	return resultChan
}


func DispatchHandlers(reqChan <-chan *Request, concurrency int) {

	for i := 0; i < concurrency; i++ {
		go handleWithTimeout(reqChan)	
		// go handle(reqChan)
	}
	
}


// normal handle
func handle(reqChan <-chan *Request) {
	var req *Request = nil
	for req = range reqChan {  // break if the reqChan is closed
		req.resultChan <- req.f(req.args ...)
	}
}


// normal handle plus timeout panic
func handleWithTimeout(reqChan <-chan *Request) {
	timeout := time.After(30 * time.Second)
	for {
		select {
			case req, ok := <- reqChan :  // 這邊需要 ok 確認 channel 是否已經關閉
				if !ok {
					reqChan = nil // prevent the channel from sending more data
					break
				}
				req.resultChan <- req.f(req.args ...)
			
			case <-timeout:  // if timeout
				err := fmt.Errorf("Reciever timeouts for waiting any incoming request")
				panic(err)  // or break the loop if you want
		}
	}
}
