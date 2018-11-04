package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"encoding/hex"
)

/**
* hmac的主要用途
* 1.用在用户登录上 ->[随机值保存在session] ( 随机值 + 密码 ) hmac运算 => （随机值 + 数据库密码） 服务器中hmac运算
*/

func main() {
	var message = "hello world"
	key := []byte("123456")
	hash := hmac.New(sha256.New, key)
	hash.Write([]byte(message))
	messageHmac := hash.Sum(nil)
	fmt.Println(hex.EncodeToString(messageHmac))

}
