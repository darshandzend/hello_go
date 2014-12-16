package main

import (
	"fmt"
	"time"
)

func main() {
	
	timer1 := time.NewTimer(time.Second * 2)

	<-timer1.C
	fmt.Println("Timer1 expired")

	timer2 := time.NewTimer(time.Second * 2)
	go func() {
		<-timer2.C
		fmt.Println("Timer2 expired")
	}()

	stop := timer2.Stop()
	if stop {
		fmt.Println("timer2 stopped.")
	}
}
