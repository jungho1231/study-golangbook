package main

import (
	"crypto/sha512"
	"fmt"
)

func main() {
	s := "Hello, world!"
	h1 := sha512.Sum512([]byte(s)) // 문자열의 SHA512 해시 값 추출
	fmt.Printf("print = %x\n", h1)

	sha := sha512.New()          // sha512 해시 인스턴스 생성
	sha.Write([]byte("Hello, ")) // 해시 인스턴스에 데이터 추가
	sha.Write([]byte("world!"))  // 해시 인스턴스에 데이터 추가
	h2 := sha.Sum(nil)           //  해스 인스턴스에 저장된 데이터의 SHA 512 해시 값 추출
	fmt.Printf("print = %x\n", h2)

}
