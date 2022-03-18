package main

import (
	"crypto/aes"
	"fmt"
)

func main() {
	key := "Hello, Key 12345" // 16 바이트
	s := "Hello, world! 12"   // 16 바이트

	block, err := aes.NewCipher([]byte(key)) // AES 대칭키 암호화 블록 생성
	if err != nil {
		fmt.Println(err)
		return
	}

	ciphertext := make([]byte, len(s))
	block.Encrypt(ciphertext, []byte(s))

	plaintext := make([]byte, len(s))
	block.Decrypt(plaintext, ciphertext)
	fmt.Println(string(plaintext))

}
