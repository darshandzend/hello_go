package main

import (
	"fmt"
	"regexp"
	"bytes"
)

func main() {
	
	matches, err:= regexp.MatchString("p([a-z]+)a", "pizza")
	fmt.Println(matches, err)

	r, _ := regexp.Compile("p([a-z]+)a")
	fmt.Println(r.MatchString("pizza"))

	fmt.Println(r.FindString("pizza pakpolya"))

	fmt.Println(r.FindStringIndex("pizza pakpolya"))

	fmt.Println(r.FindStringSubmatch("pizza pakpolya"))
	
	fmt.Println(r.FindStringSubmatchIndex("pizza pakpolya pithla"))

	fmt.Println(r.FindAllString("pizza pakpolya pithla", -1))

	fmt.Println(r.FindAllStringSubmatchIndex("pizza pakpolya pithala", -1))

	fmt.Println(r.FindAllString("pizza pakpolya pithla", 2))

	fmt.Println(r.Match([]byte("pizza")))

	r = regexp.MustCompile("p([a-z]+)a")
	fmt.Println(r)

	fmt.Println(r.ReplaceAllString("I love pizza", "<food>"))

	in := []byte("This is a pizza.")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
