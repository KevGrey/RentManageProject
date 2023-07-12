package encrypts

import (
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	str := "123456"
	md5 := Md5(str)
	fmt.Println(md5)
}

func TestEncrypt(t *testing.T) {
	plain := "3c781af02c69d319bcf1d96a073cc5cf"
	// AES 规定有3种长度的key: 16, 24, 32分别对应AES-128, AES-192, or AES-256
	key := "abcdefgehjhijkmlkjjwwoew"
	// 加密
	cipherByte, err := Encrypt(plain, key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s ==> %s\n", plain, cipherByte)
	// 解密
	plainText, err := Decrypt(cipherByte, key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s ==> %s\n", cipherByte, plainText)
}
