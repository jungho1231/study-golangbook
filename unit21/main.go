package main

import "fmt"

func main() {
	// var 배열명 [길이] 자료형

	var a [5]int // int 형이며 길이가 5인 배열 선언
	// 자바 문법에 익숙해져 있어서 new 연산자가 없어서 배열이 할당이 안된 것 같지만 이미 선언되고 0으로 초기화가 끝난 상태
	fmt.Println(a)

	a2 := [5]int{32, 29, 78, 16, 81} // 배열을 생성하면서 값을 초기화
	fmt.Println(a2, "\n")

	b := [5]int{
		32,
		29,
		78,
		16,
		81, // 여러 줄로 나열할 때는 마지막 요소에 콤마를 붙임
	}
	fmt.Println(b, "\n")

	c := [...]int{32, 29, 78, 16, 81} // 초기화할 요소가 5개이며 ...을 사용했으므로
	// 배열 크기는 5로 설정됨
	fmt.Println(c, "\n")

	d := [...]string{"Maria", "Andrew", "Jonh"} // 초기화할 요소가 3개이며 ...을 사용했으므로
	// 배열 크기는 3으로 설정됨
	fmt.Println(d, "\n")

	for i := 0; i < len(d); i++ {
		fmt.Println(d[i])
	}

	fmt.Println()

	for index, value := range d {
		fmt.Println(index, value)
	}

	fmt.Println()

	for _, value := range d {
		fmt.Println(value)
	}

	e := d

	fmt.Println(e)

	d[0] = "jung"

	fmt.Println("e = ", e)
	fmt.Println("d = ", d)
}
