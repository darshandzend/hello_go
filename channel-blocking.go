package main

import (
	"fmt"
	"time"
)

func assasination_msgs(msgs chan string, kill_status chan bool) {
	
	msgs<- "Kill the King"

	fmt.Println("Message Sent")

	select {
	case killed := <-kill_status:
		if killed {
			fmt.Println("The King is dead")
		} else {
			fmt.Println("The King has survived")
		}
	}
}

func assasinate(msgs chan string, kill_status chan bool) {
	
	time.Sleep(time.Second * 2)
	
	msg := <-msgs

	fmt.Println("Message Recieved:", msg)
	
	kill := true //or Kill(), I'm assuming the assasin won't fail

	kill_status<- kill
}


func main() {
	
	msgs := make(chan string, 1)
	kill_status := make(chan bool, 1)

	go assasination_msgs(msgs, kill_status)
	go assasinate(msgs, kill_status)

	time.Sleep(time.Second * 3)
}




