package main

import "fmt"

type rect struct {
	height, width int
}

func (r *rect) area() int {
	return r.height * r.width
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{240, 360}
	
	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	r2 := &r

	fmt.Println("area:", r2.area())
	fmt.Println("perim:", r2.perim())
}
