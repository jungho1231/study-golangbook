package main

import "fmt"

func hello() {
	fmt.Println("hello world")
}

//func sum(a, b int) (r int) {
//	r = a + b
//	return
//}
func sumAndDiff(a, b int) (int, int) {
	return a + b, a - b
}

func sum(n ...int) int {
	total := 0

	for _, value := range n {
		total += value
	}

	return total
}

func diff(a, b int) int {
	return a - b
}

func factorial(n uint64) uint64 {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func main() {

	hello()

	i := sum(1, 2)
	fmt.Println(i)

	//sum1, diff := sumAndDiff(1, 2)
	//fmt.Println(sum1, diff)

	sum2 := sum(1, 2, 3, 4, 5)
	fmt.Println(sum2)

	ints := []int{1, 2, 3, 4, 5}
	i2 := sum(ints...)
	fmt.Println(i2)

	fmt.Println(factorial(5)) // 120

	hello := sum

	fmt.Println(hello(1, 2, 3, 4, 5)) // 15

	f := map[string]func(int, int) int{
		"diff": diff,
	}

	i3 := f["diff"](3, 1)
	fmt.Println(i3) // 2

	i4 := func(a int, b int) int {
		return a + b
	}(1, 3)

	fmt.Println(i4) //4

}
