package main

import "fmt"

type Res struct {
	code string
	msg  string
}

func main() {
	res := &Res{
		"0001",
		"msg"}
	fmt.Println(res)

	res2 := &Res{
		code: "0002",
		msg: "second msg"}
	fmt.Println(res2)
}
