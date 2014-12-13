package main

import "fmt"

func incrementor() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	nextInt := incrementor()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	
	reinit := incrementor()

	fmt.Println(reinit())
	fmt.Println(reinit())
}
