package main

import "fmt"

func main() {

	i := 3
	for i > 0 {
		fmt.Println(i)
		i--
	}

	for j := 0; j < 10; j++ {
		fmt.Println(j)
	}

	for { //infinite loop
		fmt.Println("bwoy!")
	}

}
