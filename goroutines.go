package main

import "fmt"

func f(input string) {
	for i := 0; i < 4; i++ {
		fmt.Println(input, ":", i)
	}
}

func main() {
	
	//Normal call
	f("direct")

	//goroutine
	go f("goroutine")

	//anon go routine
	go func(input string) {
		for i := 0; i < 4; i++ {
			fmt.Println(input, ":", i)
		}
	}("going")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done.")
}
