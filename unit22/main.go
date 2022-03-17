package main

import "fmt"

func main() {
	// 슬라이스 동적으로 크키가 늘어난다.
	var a []int // int 형 슬라이스 선언
	fmt.Println(a)

	// int[] 가 아니라 []int
	var b []int = make([]int, 5) // 길이가 5인 슬라이스에 공간 할당
	fmt.Println(b)

	var c = make([]int, 5)
	fmt.Println(c)

	d := make([]int, 5)
	fmt.Println(d)

	// append (슬라이스, 값1,2,3...)
	a = append(a, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println(a)

	ints := []int{1, 2, 3}
	ints2 := []int{4, 5, 6}
	int3 := append(ints, ints2...)

	fmt.Println(int3)

	// 래퍼런스 타입
	i := []int{1, 2, 3}
	var j []int

	j = i

	j[0] = 9
	i[1] = 11

	fmt.Println(i)
	fmt.Println(j)

	i2 := [5]int{1, 2, 3, 4, 5}
	i3 := i2[:4] // 인덱스 4는 포함이 안됨 // 1,2,3,4
	fmt.Println(i3)

	i4 := i2[1:2] // 첫 인덱스는 포함, 마지막 인덱스는 미포함
	fmt.Println(i4)

	// 부분 슬라이스도 래퍼러스 타입
	// 슬라이스[시작 인덱스 : 끝 인덱스 : 용량]

}
