package main

import (
	"fmt"
)

func main() {
	// 7.1
	// var 변수명 자료형
	var i int
	var s string

	var age int = 10
	var name string = "Maria"

	fmt.Println(i, s, age, name)

	// 자료형을 생략할 수 있다. 이 경우에는 반드시 초기활르 해야 한다.

	// 짧은 선언 사용하기
	// 변수명 := 초기값

	age2 := 10
	name2 := "Maria"

	fmt.Println(age2, name2)

	// 7.2
	// 변수 여러개 사용하기
	// var 변수1, 변수2 자료형  = 초기값1, 초기값 2
	// var 변수1, 변수2 = 초기값1, 초기값2

	var x, y int = 30, 50 // x = 30 , y = 50

	fmt.Println(x, y)

	var age3, name3 = 10, "Jung" // age = 10, name = Jung

	fmt.Println(age3, name3)

	// 병렬할당
	var x2, y2 int
	var age4 int

	x2, y2, age3 = 10, 20, 5

	fmt.Println(x2, y2, age4)

	// var 키워드 + () 변수 여러개를 선언하고 초기화 할 수 있다

	var (
		x3, y3      = 30, 50
		age5, name5 = 10, "jung"
	)

	fmt.Println(x3, y3, age5, name5)

}
