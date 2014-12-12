package main

import "fmt"

func main() {
	
	m := make(map[string]int)

	m["k1"] = 1
	m["k2"] = 2
	fmt.Println("map: ", m)

	fmt.Println("get: ", m["k1"])

	fmt.Println("len: ", len(m))

	delete(m, "k2")
	fmt.Println("del: ", m)

	_,present := m["k2"]
	fmt.Println("prs: ", present)

	n := map[string]int{"one" : 1, "two" : 2}
	fmt.Println("dec: ", n)
}
