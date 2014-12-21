package main

import (
	"fmt"
	"encoding/json"
	"os"
)

type Response1 struct {
	Page int
	Fruits []string
}

type Response2 struct {
	Page int 		`json:"page"`
	Fruits []string	`json:"fruits"`
}

func main() {
	
	boolB, _ := json.Marshal(true)
	fmt.Println(string(boolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(3.14)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("A hat is a turban")
	fmt.Println(string(strB))

	slcD := []string{"dnyaneshwar", "tukaram", "eknath"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "banana": 4}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &Response1{
		Page: 1,
		Fruits: []string{"apple", "banana", "chiku"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &Response2{
		Page: 2,
		Fruits: []string{"apple", "banana", "chiku"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	byt := []byte(`{"num":5.23, "strs":["a", "b"]}`)

	var dat map[string]interface{}


	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64)
	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	str := `{"page": 1, "fruits": ["banana", "apple", "chicku"]}`
	res := &Response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "banana": 3}
	enc.Encode(d)
}
