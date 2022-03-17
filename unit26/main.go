package main

import (
	"fmt"
	"os"
)

func hello() {
	fmt.Println("hello")
}

func world() {
	fmt.Println("world")
}

func helloWorld() {
	defer func() {
		fmt.Println("world")
	}()

	func() {
		fmt.Println("hello")
	}()
}

func ReadHello() {
	file, err := os.Open("hello.txt")

	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	buf := make([]byte, 100)

	_, err = file.Read(buf)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(buf))
}

func main() {

	// defer는 특정 함수를 현재 함수가 끝나기 직전에 실행한다. try finally 구문과 비슷
	//defer world()
	//hello()
	//hello()
	//hello()

	//helloWorld()

	ReadHello()
}
