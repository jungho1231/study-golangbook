package main

import (
	"fmt"
	"os"
)

func main() {
	var num1 int
	var num2 float32
	var s string

	file1, _ := os.Open("hello1.txt")
	defer file1.Close()
	n, _ := fmt.Fscan(file1, &num1, &num2, &s)

	fmt.Println("입력 개수 ", n)
	fmt.Println(num1, num2, s)

	file2, _ := os.Open("hello2.txt")
	defer file2.Close()
	fmt.Fscanln(file2, &num1, &num2, &s)

	fmt.Println("입력 개수 ", n)
	fmt.Println(num1, num2, s)

	file3, _ := os.Open("hello3.txt")               // hello3.txt 파일 열기
	defer file3.Close()                             // main 함수가 끝나기 직전에 파일을 닫음
	fmt.Fscanf(file3, "%d,%f,%s", &num1, &num2, &s) // 파일을 읽은 뒤 문자열에서
	// 형식을 지정하여 입력을 받음
	fmt.Println("입력 개수:", n) // 입력 개수: 3
	fmt.Println(num1, num2, s)
}
