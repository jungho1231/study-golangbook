package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// 정수 생략

	// 소수
	var a float64 = 10.0

	for i := 0; i < 10; i++ {
		a = a - 0.1
	}

	fmt.Println(a)        // 반올림 오차 발생
	fmt.Println(a == 9.0) // false

	// todo el 변수가 작동안함.
	// https://golangbyexample.com/comparing-floating-point-numbers-go/ 링크로 예제를 대신함
	//const epsilon = el - 14

	// 복소수 생략

	// 바이트
	var num1 byte = 10   // 10 진수 저장
	var num2 byte = 0x32 // 16 진수 저장
	var num3 byte = 'a'  // 문자 저장

	fmt.Println(num1, num2, num3)

	/*
		var num1 byte = "a"  // 컴파일 에러
		var num2 byte = 'ab' // 컴파일 에러
		var num3 byte = "ab" // 컴파일 에러
	*/

	// 룬 (UTF - 8)

	var r1 rune = '한'
	var r2 rune = '\ud55c'     // 한
	var r3 rune = '\U0000d55c' // 한

	fmt.Println(r1, r2, r3)
	fmt.Println("rune 한 = ", string(r1), string(r2), string(r3))

	// 숫자 연산하기 생략
	// 자료형이 다른 숫자 연산을 할 경우 자료형에 따라 누락되는 값이 있을 수 있다.

	// 오버 플로우, 언더 플로우 생략
	// 각 자료형마다 저장할 수 있는 최대 크기를 넘어서면 오버 플로우, 최소 크기보다 작아지면 언더플로우 발생

	var num4 int8 = 1
	var num5 int16 = 1
	var num6 int32 = 1
	var num7 int64 = 1

	fmt.Println(unsafe.Sizeof(num4)) // 1
	fmt.Println(unsafe.Sizeof(num5)) // 2
	fmt.Println(unsafe.Sizeof(num6)) // 4
	fmt.Println(unsafe.Sizeof(num7)) // 4
}
