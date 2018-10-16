package main

import (
	"fmt"
)

const name  = "xioaliao"

func main() {
	var (
		a int
		b string
		c []int
		d [3]int
		e bool = true
		g func() string
		h struct{}
	)

	a, b, f := 1, "2", 3
	f, b, a = a, b, f;
	g = func() string{
		return "hi"
	}

	fmt.Printf("%T,%v\n", name, name)
	fmt.Printf("%T,%v\n", a, a)
	fmt.Printf("%T,%v\n", b, b)
	fmt.Printf("%T,%v\n", c, c)
	fmt.Printf("%T,%v\n", d, d)
	fmt.Printf("%T,%v\n", e, e)
	fmt.Printf("%T,%v\n", f, f)
	fmt.Printf("%T,%v\n", g, g)
	fmt.Printf("%T,%v\n", h, h)


}
