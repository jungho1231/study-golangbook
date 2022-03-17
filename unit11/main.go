package main

import "fmt"

func main() {
	// const 상수명 자료형 = 초기값
	// go에서는 상수를 대문자로 쓰는 관례는 없는 것 같다
	const age int = 11
	const name string = "jung"

	/*
		const score int  // 컴파일 에러
		age = 11 // 컴파일 에러
		name = "jungho2" //컴파일 에러
	*/

	// const 상수명1 상수명2 자료형 = 초기값1, 초기값2
	// const 상수명1, 상수명 2 = 초기값1 초기값2

	const (
		x, y        int = 30, 50
		age2, name2     = 10, "jung"
	)

	fmt.Println(x, y, age2, name2)

}
