package main

import (
	"fmt"
	"crypto/x509"
)

func testConsts(){
	fmt.Println("x509 RSA  is :",int(x509.RSA)) //1
	fmt.Println("x509 DSA  is :",int(x509.DSA)) //2
	fmt.Println("x509 ECDSA  is :",int(x509.ECDSA)) //3

	//KeyUsage 秘钥操作
	fmt.Println("KeyUsageDigitalSignature  is",int(x509.KeyUsageDigitalSignature)) //1
	fmt.Println("KeyUsageContentCommitment is",int(x509.KeyUsageContentCommitment)) //2
	fmt.Println("KeyUsageKeyEncipherment is",int(x509.KeyUsageEncipherOnly)) //4
	fmt.Println("KeyUsageDataEncipherment is",int(x509.KeyUsageDataEncipherment)) //8
	fmt.Println("KeyUsageKeyAgreement is",int(x509.KeyUsageKeyAgreement)) //16
	fmt.Println("KeyUsageCertSign is",int(x509.KeyUsageCertSign)) //32
	fmt.Println("KeyUsageCRLSign is",int(x509.KeyUsageCRLSign)) //64
	fmt.Println("KeyUsageEncipherOnly is",int(x509.KeyUsageEncipherOnly)) //128
	fmt.Println("KeyUsageDecipherOnly is",int(x509.KeyUsageDecipherOnly)) //256

	//ExtKeyUsage 扩展秘钥操作
	fmt.Println("ExtKeyUsageServerAuth is",int(x509.ExtKeyUsageServerAuth)) //1
	fmt.Println("ExtKeyUsageClientAuth is",int(x509.ExtKeyUsageClientAuth)) //2
	fmt.Println("ExtKeyUsageCodeSigning is",int(x509.ExtKeyUsageCodeSigning)) //3
	fmt.Println("ExtKeyUsageEmailProtection is",int(x509.ExtKeyUsageEmailProtection)) //4
	fmt.Println("ExtKeyUsageIPSECEndSystem is",int(x509.ExtKeyUsageIPSECEndSystem)) //5
	fmt.Println("ExtKeyUsageIPSECTunnel is",int(x509.ExtKeyUsageIPSECTunnel)) //6
	fmt.Println("ExtKeyUsageIPSECUser is",int(x509.ExtKeyUsageIPSECUser)) //7
	fmt.Println("ExtKeyUsageTimeStamping is",int(x509.ExtKeyUsageTimeStamping)) //8
	fmt.Println("ExtKeyUsageOCSPSigning is",int(x509.ExtKeyUsageOCSPSigning)) //9
	fmt.Println("ExtKeyUsageMicrosoftServerGatedCrypto is",int(x509.ExtKeyUsageMicrosoftServerGatedCrypto)) //10
	fmt.Println("ExtKeyUsageNetscapeServerGatedCrypto is",int(x509.ExtKeyUsageNetscapeServerGatedCrypto)) //1
}
func main() {
	testConsts()
}
