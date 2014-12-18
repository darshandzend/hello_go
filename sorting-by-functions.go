package main

import (
	"fmt"
	"sort"
)

type person struct {
	name string
	age  int
}

type ByAge []person

func (b ByAge) Len() int {
	return len(b)
}

func (b ByAge) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByAge) Less(i, j int) bool {
	return b[i].age < b[j].age
}

func main() {

	people := []person{person{"Ravi", 29}, person{"Som", 33}, person{"Mangal", 12}}
	sort.Sort(ByAge(people))
	fmt.Println(people)
}
