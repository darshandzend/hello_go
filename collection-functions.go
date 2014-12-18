package main

import (
	"fmt"
	"strings"
)

func Index(vs []string, t string) int {
	for i, s := range vs {
		if s == t {
			return i
		}
	}
	return -1
}

func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

func Any(vs []string, f func(string) bool) bool {
	for _, s := range vs {
		if f(s) {
			return true
		}
	}
	return false
}

func All(vs []string, f func(string) bool) bool {
	for _, s := range vs {
		if !f(s) {
			return false
		}
	}
	return true
}

func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, s := range vs {
		if f(s) {
			vsf = append(vsf, s)
		}
	}
	return vsf
}

func Map(vs []string, f func(string) string) []string {
	vsf := make([]string, len(vs))
	for i, s := range vs {
		vsf[i] = f(s)
	}
	return vsf
}

func main() {

	var strs = []string{"Varan", "Bhat", "Bhaaji", "Poli"}

	fmt.Println(Index(strs, "Bhat"))

	fmt.Println(Include(strs, "Dahi"))

	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "B")
	}))

	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "B")
	}))

	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "a")
	}))

	fmt.Println(Map(strs, strings.ToUpper))

}
