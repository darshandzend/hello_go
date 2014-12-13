package main

import "fmt"

func sum(nums ...int) int {
	
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}

	return total
}

func main() {
	
	fmt.Println(sum(1, 2))
	
	fmt.Println(sum(1, 3, 5 , 383))

	nums := []int{23, 45, 22}
	fmt.Println(sum(nums...))
}
