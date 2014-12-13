package main

import "fmt"

func ping(pings chan<- string, msg string) {
	fmt.Println("<Ping>")
	pings <- msg
	fmt.Println("</Ping>")
}

func pong(pings <-chan string, pongs chan<- string) {
	fmt.Println("<Pong>")
	msg := <-pings
	pongs <- msg
	fmt.Println("</Pong>")
}

func main() {
	
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "Here comes the message!")
	pong(pings, pongs)

	fmt.Println(<-pongs)
}
