package main

import "fmt"

func main() {
	const (
		A = 1
		B = 'a'
		C = iota
		D
		E
	)
	fmt.Printf("%v,%c\n", A, A)
	fmt.Printf("%v,%c\n", B, B)
	fmt.Printf("%v,%c\n", C, C)
	fmt.Printf("%v,%c\n", D, D)
	fmt.Printf("%v,%c\n", E, E)


}
