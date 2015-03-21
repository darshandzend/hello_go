package main

import "fmt"
import "flag"

func main() {

	wordPtr := flag.String("word", "foo", "A string")

	numbPtr := flag.Int("numb", 42, "An Int")

	boolPtr := flag.Bool("fork", false, "A bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string bar")

	flag.Parse()

	p := fmt.Println

	p("Word:", *wordPtr)
	p("Numb:", *numbPtr)
	p("Bool:", *boolPtr)
	p("Svar:", svar)
	p("tail:", flag.Args())
}
