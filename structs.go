package main

import "fmt"

type person struct {
	name string
	age int
}

func main() {
	
	fmt.Println(person{"Michael Scott", 40})

	fmt.Println(person{name: "Gary Oldman", age: 105})

	//fmt.Println(person{"Linda Clark"})			//generates error

	fmt.Println(person{name: "Vivek Anand"})

	s := person{"Thomp Butts", 43}
	fmt.Println(s.name, s.age)

	s2 := &s
	fmt.Println(s2.name)

	s2.name = "Chidiya Jack"
	fmt.Println(s2.name)
}
