package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	var s7 string = `안녕하세요
Hello, world!`

	fmt.Println(s7)

	var s1 string = "한글"
	var s2 string = "Hello"

	fmt.Println(len(s1)) // 6  UTF-8
	fmt.Println(len(s2)) // 5 문자열 길이 그대로

	fmt.Println(utf8.RuneCountInString(s1)) // 2

	// 문자열 연산하기

	var s3 string = "한글"
	var s4 string = "Go"
	fmt.Println(s3[0])
	fmt.Println(s4[0])

	// 문자열 수정하기 생략
	// 문자열 선언 후 다른 문자열로 대입은 가능하지만 만주열의 내용을 수정하는 경우는 컴파일 에러가 난다.

}
