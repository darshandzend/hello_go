package main

import (
	"fmt"
	"math"
	"os"
)

type point struct {
	x, y int
}

func main() {
	
	p := point{3, 5}
	fmt.Printf("%v\n", p)

	fmt.Printf("%+v\n", p)

	fmt.Printf("%#v\n", p)

	fmt.Printf("%T\n", p)

	fmt.Printf("%t\n", true)

	fmt.Printf("%d\n", 124)

	fmt.Printf("%b\n", 2015)

	fmt.Printf("%c\n", 2015)

	fmt.Printf("%x\n", 2015)

	fmt.Printf("%f\n", math.Pi)

	fmt.Printf("%e\n", math.Pi)
	fmt.Printf("%E\n", math.Pi)

	fmt.Printf("%s\n", "\"Rights\"")

	fmt.Printf("%x\n", "hex this")

	fmt.Printf("%p\n", &p)

	fmt.Printf("|%6d|%4d|\n", 2223, 23)

	fmt.Printf("|%6.2f|%4.3f|\n", 2.4455, 2.4)

	fmt.Printf("|%-6.2f|%-4.3f|\n", 2.3344, 2.9)

	fmt.Printf("%20s : %-20s\n", "Directed by", "Goldy Fishtank")

	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os. Stderr, "an %s\n", "error")
}
