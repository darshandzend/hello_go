package main

import "fmt"
import "strconv"

func main() {
	
	f, _ := strconv.ParseFloat("1.2343", 64)
	fmt.Println(f)

	i, _ := strconv.ParseInt("2343", 0, 64)
	fmt.Println(i)

	d, _ := strconv.ParseInt("0xabab", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseUint("234", 0, 64)
	fmt.Println(u)

	i2, _ := strconv.Atoi("2343")
	fmt.Println(i2)

	a := strconv.Itoa(2343)
	fmt.Println(a)

	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
