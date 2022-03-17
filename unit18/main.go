package main

import "fmt"

// goto문 자체가 나쁜 것은 아니나, 오용으로 인해 스파게티 코드가 만들어질 가능성이 높다.
// 안전 장치가 없기 때문에 실수가 나올 가능성이 크다.
// goto문을 써야할 필요가 분명하지 않다면 쓰지 않는 것이 좋겠다.
func main() {
	// goto
	var a int = 1

	if a == 1 {
		goto ERROR1 // 에러 상황이면  ERROR1 로 이동
		fmt.Println("after Error1")
	}

	if a == 2 {
		goto ERROR2
	}

	fmt.Println(a)
	fmt.Println("Success")

ERROR1:
	fmt.Println("Error 1")
	return

ERROR2:
	fmt.Println("Error 2")
	return
}
