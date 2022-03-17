package main

import "fmt"

func main() {

	// for 조건식 { 변화식 }
	// while 문 처럼 동작한다

	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	//for {
	//	fmt.Println("hello world ") // 무한 루프 생성
	//}

	// 반북문에서 변수 어려 개 사용하기

	for i, j := 0, 0; i < 10; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}

	// 병렬 할당만 사용 가능
	//for i, j := 0, 0; i < 10; i++, j+=2 { // 컴파일 에러. syntax error: unexpected comma, expecting {
	//	fmt.Println(i, j)
	//}
}
