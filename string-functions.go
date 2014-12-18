package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {

	p("Contains:", s.Contains("test", "es"))
	p("Count:", s.Count("Karnataka", "a"))
	p("HasPrefix:", s.HasPrefix("Doctor Who", "Doctor"))
	p("HasSuffix:", s.HasSuffix("Ivan Junior", "Junior"))
	p("Index:", s.Index("Shree Ganeshay Namah:", "Ganesh"))
	p("Join:", s.Join([]string{"Maha", "Rashtra"}, "-"))
	p("Repeat:", s.Repeat("a", 5))
	p("Replace:", s.Replace("Dharave", "Dha", "Ma", -1))
	p("Replace:", s.Replace("Dhama yanna Dharave", "Dha", "Ma", 1))
	p("Split:", s.Split("Rock-and-Roll", "-"))
	p("ToLower:", s.ToLower("TEST"))
	p("ToUpper:", s.ToUpper("test"))
	p()

	p("Len:", len("Hello"))
	p("Char:", "Hello"[3])
}
