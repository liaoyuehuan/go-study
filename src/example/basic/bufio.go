package main

import (
	"strings"
	"bufio"
	"fmt"
)

func main() {
	var rd *strings.Reader = strings.NewReader("hello world")
	var brd *bufio.Reader = bufio.NewReader(rd)
	str,err := brd.ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Println(str)
}
