package main

import "fmt"

func main() {
	arr := [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l"}
	s1 := arr[3:6]
	s2 := arr[4:8]
	fmt.Printf("arr type : %T", arr)
	fmt.Println(arr, s1, s2)
	fmt.Println(s1[5:9], s2[5:8])

	s := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l"}
	fmt.Printf("s type : %s", s)
	s = s[:len(s)-5]
	fmt.Println(s[0:12]) // "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l"
}
