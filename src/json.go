package main

import (
	"fmt"
	"encoding/json"
	"os"
)

type Respone1 struct {
	code int
	data []string
}

func main() {
	fmt.Println("json")

	var res1 *Respone1 = &Respone1{
		code: 1,
		data: []string{"hi", "ha"}}

	fmt.Println(*res1)

	var a map[string]interface{}
	jsonStr := `{"num":6.13,"strs":["a","b"]}`
	err := json.Unmarshal([]byte(jsonStr),&a)
	if err != nil {
		panic("json unmarshal error")
	}
	fmt.Println(a)

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
