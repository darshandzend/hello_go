package main

import (
	"fmt"
	"time"
)

func main() {

	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(time.Millisecond * 200)
	for r := range requests {
		<-limiter
		fmt.Println("Request", r, time.Now())
	}

	/* Bursty Limiting */

	const REQUESTS_LIMIT int = 10

	burstyRequests := make(chan int, REQUESTS_LIMIT)
	for j := 1; j <= REQUESTS_LIMIT; j++ {
		burstyRequests <- j
	}

	burstyLimiter := make(chan time.Time, 3)
	for k := 1; k <= 3; k++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	for l := 1; l <= REQUESTS_LIMIT; l++ {
		<-burstyLimiter
		fmt.Println("Request", l, time.Now())
	}
}
