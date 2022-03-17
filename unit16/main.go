package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//file, err := ioutil.ReadFile("./hello.txt")
	//
	//if err == nil {
	//	fmt.Printf("%s", file)
	//}

	if b, err := ioutil.ReadFile("./hello.txt"); err == nil {
		fmt.Printf("%s", b)
	}

}
