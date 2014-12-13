package main

import "fmt"

func plus(a int, b int) int {
	
	return a + b
}

func main() {
	
	res := plus(3, 5)
	fmt.Println("res: ", res)
}

