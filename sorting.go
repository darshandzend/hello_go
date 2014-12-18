package main

import (
	"fmt"
	"sort"
)

func main() {

	strs := []string{"foo", "bar", "baz"}
	sort.Strings(strs)
	fmt.Println("Sorted:", strs)

	ints := []int{23, 32, 84}
	sort.Ints(ints)
	fmt.Println("Sorted:", ints)

	fmt.Println("Is sorted:", sort.IntsAreSorted(ints))
}
