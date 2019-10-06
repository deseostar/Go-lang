package main

import (
	"fmt"
)

// b := 2 //error

func main() {
	var a int = 1
	// var a, b int
	var c = 1
	b := 2
	var d int
	var e string
	const f int = 3 //값 변경 안됨

	fmt.Println(a, b, c, d, e)
}