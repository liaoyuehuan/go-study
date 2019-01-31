package main

import (
	"crypto/rsa"
	"crypto/rand"
	"fmt"
	"crypto/sha256"
	"crypto"
	"encoding/base64"
)

func printPrivateKey(privateKet *rsa.PrivateKey) {
	fmt.Println("私有的指数 D :", privateKet.D)
	fmt.Println("N的素因子 Primes : ", privateKet.Primes)
	fmt.Println("Precomputed :", privateKet.Precomputed)
}

func testPSS() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	plaintText := []byte("hello pss")
	hash := sha256.New()
	hash.Write(plaintText)
	hashed := hash.Sum(nil)
	//signPss
	signature, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, hashed, &rsa.PSSOptions{
		SaltLength: rsa.PSSSaltLengthAuto})
	if err != nil {
		panic(err)
	}
	fmt.Println("signature len : ", len(signature))
	fmt.Println("signature data :", base64.StdEncoding.EncodeToString(signature))

	//verify
	err = rsa.VerifyPSS(&privateKey.PublicKey, crypto.SHA256, hashed, signature, &rsa.PSSOptions{
		SaltLength: rsa.PSSSaltLengthAuto})
	if err != nil {
		panic(err)
	} else {
		fmt.Print("verify Success")
	}



}

func main() {
	//生成一个私钥
	testPSS()
}
