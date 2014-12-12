package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	
	fmt.Println(s)

	const i = 40983483

	const f = 2e18 / i
	fmt.Println(f)

	//fmt.Println(int64(f))

	fmt.Println(math.Sin(f))

}
