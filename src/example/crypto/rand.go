package main

import (
	"crypto/rand"
	"math/big"
	"fmt"
)

//rand.Int(rand io.Reader,max *Big.Int) *big.Int,err  返回[0,max)区间的随机值
func testRandInt()  {
	randInt,err :=  rand.Int(rand.Reader,big.NewInt(600000))
	if err != nil {
		panic(err)
	}
	fmt.Printf("rand int is %v \n",randInt)
}

func testRandRead()  {
	var randByte = make([]byte,20)
	randLen,err := rand.Read(randByte)
	if err != nil {
		panic(err)
	}
	fmt.Println("randLen is :",randLen)
	fmt.Println("randByte is :",randByte)
}

func main(){
	testRandInt()
	testRandRead()
}
