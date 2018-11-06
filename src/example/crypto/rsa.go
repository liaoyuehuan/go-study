package main

import (
	"crypto/rsa"
	"crypto/rand"
	"fmt"
)

func printPrivateKey(privateKet *rsa.PrivateKey)  {
	fmt.Println("私有的指数 D :",privateKet.D)
	fmt.Println("N的素因子 Primes : ",privateKet.Primes)
	fmt.Println("Precomputed :",privateKet.Precomputed)
}

func main(){
	//生成一个私钥
	privateKey,err := rsa.GenerateKey(rand.Reader,1024)
	if err != nil {
		panic(err)
	}
	printPrivateKey(privateKey)
}