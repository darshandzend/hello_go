package main

import (
	"fmt"
	"time"
)

func main() {
	
	ticker := time.NewTicker(time.Millisecond * 500)

	go func() {
		for t := range ticker.C {
			fmt.Println("Ticker at", t)
		}
	}()

	time.Sleep(time.Millisecond * 1500)
	ticker.Stop()
	fmt.Println("Ticker Stopped.")
}
