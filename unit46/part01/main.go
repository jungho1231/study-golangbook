package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Hello, world!", "wo"))  // true
	fmt.Println(strings.Contains("Hello, world!", "w o")) // false
	fmt.Println(strings.Contains("Hello, world!", "ow"))  // false

	fmt.Println(strings.ContainsAny("Hello, world!", "wo"))  // true
	fmt.Println(strings.ContainsAny("Hello, world!", "w o")) // true
	fmt.Println(strings.ContainsAny("Hello, world!", "ow"))  // true

	fmt.Println(strings.Count("Hello Helium", "He")) // 2}

	var r rune
	r = '하'

	fmt.Println(strings.ContainsRune("안녕하세요", r)) // true

	s := "Hello, world!"
	prefix := "He"
	fmt.Println(strings.HasPrefix(s, prefix)) // true
	fmt.Println(strings.HasSuffix(s, "rld!")) // true

}
