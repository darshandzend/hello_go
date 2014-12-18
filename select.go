package main

import "fmt"
import "time"

func main() {

	ch1 := make(chan string, 1)
	ch2 := make(chan string, 2)

	go func() {
		time.Sleep(time.Second * 1)
		ch1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		ch2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("recieved:", msg1)
		case msg2 := <-ch2:
			fmt.Println("recieved:", msg2)
			//default:
			//	fmt.Println("Nothing recieved yet")
		}
	}
}
