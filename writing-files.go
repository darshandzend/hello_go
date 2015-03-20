package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	d1 := []byte("Hello\ngo\n")
	err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	f, err := os.Create("/tmp/dat2")
	check(err)

	defer f.Close()

	d2 := []byte{100, 101, 102, 103, 105}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("%d bytes written\n", n2)

	n3, err := f.WriteString("I'm so sleepy\n")
	check(err)
	fmt.Printf("%d bytes written\n", n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("I'm effin' tired\n")
	check(err)
	fmt.Printf("%d bytes written\n", n4)

	w.Flush()

}
