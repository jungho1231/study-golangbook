package main

import "fmt"

func main() {
	//const (
	//	Sunday       = 0
	//	Monday       = 1
	//	Tuesday      = 2
	//	Wednesday    = 3
	//	Thursday     = 4
	//	Friday       = 5
	//	Saturday     = 6
	//	numberOfDays = 7
	//)

	const (
		Sunday       = iota // 0
		Monday              // 1
		Tuesday             // 2
		Wednesday           // 3
		Thursday            // 4
		Friday              // 5
		Saturday            // 6
		numberOfDays        // 7
	)

	fmt.Println(Thursday)     // 4
	fmt.Println(numberOfDays) // 7
	fmt.Println(Friday)

	fmt.Println("=============================")

	// 쉬프트 연산자로 << 왼쪽으로 밀면 값이 2배씩 증가한다. 반대로 밀면 2배씩 감소
	const (
		a = 1 << iota // a == 1 (1 << 0)
		b = 1 << iota // b == 2 (1 << 1)
		c = 1 << iota // c == 4 (1 << 2)
		d = 1 << iota // d == 8 (1 << 3)
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	//const (
	//	a = iota * 30 // a == 0 (0 * 30)
	//	b = iota * 30 // b == 30 (1 * 30)
	//	c = iota * 30 // c == 60 (2 * 30)
	//	d = iota * 30 // d == 90 (3 * 30)
	//)

	const (
		bit0, mask0 = 1 << iota, 1<<iota - 1 // bit0 == 1, mask0 == 0
		bit1, mask1                          // bit1 == 2, mask1 == 1
		_, _                                 // iota == 2를 생략하여 bit2와 mask2 생략
		bit3, mask3                          // bit3 == 8, mask3 == 7
	)
}
