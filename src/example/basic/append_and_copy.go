package main

import (
	"fmt"
)

func main() {
	var numbers []int
	fmt.Println("1.-------------")
	printSlice("numbers", numbers) //nil
	numbers = append(numbers, 0)
	printSlice("numbers", numbers) //[0]
	ints := []int{1, 2, 3, 4, 5, 6, 7}
	numbers = append(numbers, ints...)
	printSlice("numbers", numbers) //[0,1,2,3,4,5,6,7] cap 8
	numbers_ := numbers
	numbers = append(numbers, 8)
	numbers_[0] = 100
	printSlice("numbers", numbers) //[0,1,2,3,4,5,6,7,8] cap 16
	printSlice("numbers_",numbers_)
	fmt.Println("2.-------------")
	numbers2 := make([]int, len(numbers), cap(numbers)*2)
	copy(numbers2, numbers)
	printSlice("number2", numbers2) //[0,1,2,3,4,5,6,7,8] cap 32

	numbers2 = append(numbers2, []int{9, 10, 11}...)
	printSlice("number2", numbers2) //[0,1,2,3,4,5,6,7,8,9,10,11] cap 32 地址不变

	
}

func printSlice(name string, x []int) {
	fmt.Printf("%s\t%p\t%d\t%d\t%v\n", name, x, len(x), cap(x), x)
}
