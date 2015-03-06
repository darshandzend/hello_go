package main

import "fmt"
import "math/rand"

func main() {
	p := fmt.Println

	p(rand.Intn(5))

	s := rand.NewSource(20)
	r := rand.New(s)

	p(r.Intn(342))
	p(r.Intn(342))

	s2 := rand.NewSource(20)
	r2 := rand.New(s2)

	p(r2.Intn(342))
	p(r2.Intn(342))
}
