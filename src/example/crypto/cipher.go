package main

import (
	"crypto/aes"
	"io"
	"crypto/rand"
	"crypto/cipher"
	"fmt"
	"encoding/base64"
	"strings"
)

func testCfb() {
	//key only 16 24 32
	key := []byte("1234567812345678");

	//plainText
	plainText := []byte("hello world")

	block, err := aes.NewCipher(key)

	//encrypt
	if err != nil {
		panic(err)
	}
	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize ]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	fmt.Println(cap(iv))
	cfbCipher := cipher.NewCFBEncrypter(block, iv)
	cfbCipher.XORKeyStream(cipherText[aes.BlockSize:], plainText)
	cipherBase64 := base64.StdEncoding.EncodeToString(cipherText)
	fmt.Println("iv : ", string(iv))
	fmt.Println("cipherText : ", base64.StdEncoding.EncodeToString(cipherText[aes.BlockSize:]))

	//decrypt
	cipherText, err = base64.StdEncoding.DecodeString(cipherBase64)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	iv = cipherText[:aes.BlockSize]
	cfbCipher = cipher.NewCFBEncrypter(block, iv)
	cfbCipher.XORKeyStream(plainText, cipherText[aes.BlockSize:])
	fmt.Println("plainText : ", string(plainText))
}

func testCBC() {
	key := []byte("1234567887654321")
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	plainText := []byte("hello cbc")
	plainText = []byte(pkcs5Pad(string(plainText), aes.BlockSize))
	if len(plainText)%aes.BlockSize != 0 {
		panic("plaintext is not a multiple of the block size")
	}

	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	//encrypt
	cbc := cipher.NewCBCEncrypter(block, iv)
	cbc.CryptBlocks(cipherText[aes.BlockSize:], plainText)
	fmt.Println("cipher text:", base64.StdEncoding.EncodeToString(cipherText[aes.BlockSize:]))

	//decrypt
	dePlainText := make([]byte, 20)
	iv = cipherText[:block.BlockSize()]
	deCbc := cipher.NewCBCDecrypter(block, iv)
	deCbc.CryptBlocks(dePlainText, cipherText[aes.BlockSize:])
	pad := int(dePlainText[len(dePlainText)-1])
	fmt.Println("plaint text:", string(dePlainText[:len(dePlainText)-pad]), )

}

func pkcs5Pad(input string, blockSize int) string {
	size := len(input)
	pad := blockSize - size%blockSize
	input = input + strings.Repeat(string(pad), pad)
	return input
}

func main() {
	testCBC()
}
