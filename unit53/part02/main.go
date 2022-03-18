package main

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

func main() {

	// rand.Reader - 보안 관련된 랜덤 숫자 생성
	// secure random number generator.
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048) // 개인 키와 공개 키
	if err != nil {
		fmt.Println(err)
		return
	}

	publicKey := privateKey.PublicKey // 개인 키 변수 안에 공개 키가 들어있음

	message := "안녕하세요. Go 언어"
	hash := md5.New()           // 해시 인스턴스 생성
	hash.Write([]byte(message)) // 해시 인스턴스에 문자열 추가
	digest := hash.Sum(nil)     // 문자열의 MD5 해시 값 추출

	// 번역기 : Hash는 다른 코드에서 구현된 암호화 해시 함수를 식별합니다
	// 암호화 식별자 ?
	var h1 crypto.Hash

	signature, err := rsa.SignPKCS1v15(
		rand.Reader,
		privateKey, // 개인 키
		h1,
		digest, // MD5 해시 값
	)

	var h2 crypto.Hash

	err = rsa.VerifyPKCS1v15( // 공개 키로 서명 검증
		&publicKey, // 공개 키
		h2,
		digest,    // MD5 해시 값
		signature, // 서명 값
	)

	if err != nil {
		fmt.Println("검증 실패 ")
	} else {
		fmt.Println("검증 성공 ")
	}

}
