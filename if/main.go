package main

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func main() {
	var b bool = true

	if b {
		fmt.Println("b is true")
	}

	if num := add(1,2); num == 3 {
		fmt.Println("num is 3")
	} else if num == 5 {
		
	} else {
		
	}
}