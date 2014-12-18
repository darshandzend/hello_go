package main

import "os"

func main() {

	panic("A Problem")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
