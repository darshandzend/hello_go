package main

import "fmt"

func main() {

	jobs := make(chan int, 5)
	done := make(chan bool, 1)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("Recieved Job", j)
			} else {
				done <- true
			}
		}
	}()

	for i := 0; i < 3; i++ {
		jobs <- i
		fmt.Println("Sent Job", i)
	}

	close(jobs)
	fmt.Println("Sent all Jobs.")

	<-done
}
