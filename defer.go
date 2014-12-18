package main

import (
	"fmt"
	"os"
)

func main() {

	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(path string) *os.File {

	fmt.Println("Creating...")

	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	return f
}

func writeFile(f *os.File) {

	fmt.Println("Writing...")

	fmt.Fprintln(f, "If the Gods were human, they would eat Rice.")
}

func closeFile(f *os.File) {

	fmt.Println("Closing")

	f.Close()
}
