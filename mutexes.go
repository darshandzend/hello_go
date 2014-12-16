package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	
	state := make(map[int]int)

	var ops int64 = 0

	var mutex = &sync.Mutex{}

	
	//create 50 concurrent readers
	for i := 0; i < 50; i++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)

				runtime.Gosched()
			}
		}()
	}


	//create 10 concurrent writers
	for j := 0; j < 10; j++ {
		go func() {
			for {
				key := rand.Intn(5)
				value := rand.Intn(100)
				mutex.Lock()
				state[key] = value
				mutex.Unlock()

				atomic.AddInt64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)

	mutex.Lock()
	fmt.Println("state:",state)
	mutex.Unlock()
}
