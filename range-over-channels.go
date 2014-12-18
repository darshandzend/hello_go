package main

import "fmt"

func main() {

	ch := make(chan string, 3)
	ch <- "Ek"
	ch <- "Don"

	close(ch)

	for str := range ch {
		fmt.Println("Recieved", str)
	}
}
