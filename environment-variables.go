package main

import "os"
import "strings"
import "fmt"

func main() {
	os.Setenv("FOO", "I'm robot 234")

	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	os.Setenv("BAR", "where them code at?")

	fmt.Println()
	for _, env := range os.Environ() {
		pair := strings.Split(env, "=")
		fmt.Println(pair[0], ":", pair[1])
	}
}
