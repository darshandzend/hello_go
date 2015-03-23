package main

import "fmt"
import "os"

func main() {

	defer fmt.Println("Print this if go isn't awesome!!")

	os.Exit(3)
}
