package main

import (
	"fmt"
	"reflect"
)

func main() {
	
	var a string = "initial"
	fmt.Println(a)

	var b, c int = 1, 3
	fmt.Println(b + c)

	var d int
	fmt.Println(d)

	var e = 0
	fmt.Println(reflect.TypeOf(e))

	f := "fire"
	fmt.Println(f)

}
