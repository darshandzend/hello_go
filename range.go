package main

import "fmt"

func main() {

	nums := []int{2, 3, 5}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum: ", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("ind: ", i)
		}
	}

	kvs := map[string]int{"one": 1, "two": 2, "three": 3}
	for k, v := range kvs {
		fmt.Printf("%s -> %d\n", k, v)
	}

	for i, c := range "darshan" {
		fmt.Println(i, c)
	}
}
