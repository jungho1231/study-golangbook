package main

func Calc() func(x int) int {
	a, b := 3, 5

	return func(x int) int {
		return a*x + b
	}
}

func main() {

}
