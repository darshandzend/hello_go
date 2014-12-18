package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 2)
		ch1 <- "Relese the Kraken!"
	}()

	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case <-time.After(time.Second * 3):
		fmt.Println("Timeout")
	}
}
