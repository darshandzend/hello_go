package main

import "fmt"

func vals() (int, int) {

	return 3, 5
}

func main() {

	a, b := vals()
	fmt.Printf("a and b: %d, %d\n", a, b)
}
