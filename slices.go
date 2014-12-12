package main

import "fmt"

func main() {

	s := make([]string, 3)
	fmt.Println("emp: ", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("get: ", s)

	fmt.Println("len: ", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("app: ", s)

	c := make([]string, len(s))
	copy(c,s)
	fmt.Println("cpy: ", c)

	l := s[2:5]
	fmt.Println("sl1: ", l)

	t := []string{"one", "two", "three"}
	fmt.Println("dcl: ", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerlen := i + 1
		twoD[i] = make([]int, innerlen)
		for j := 0; j < innerlen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d : ", twoD)

}
