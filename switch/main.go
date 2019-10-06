package main

import (
	"fmt"
)

func main() {
	// s := "abc"
	s := "ab"

	switch s {
	case "abc":
		fmt.Println("s is abc")
		// fallthrough 값이 맞으면 넘어감?
	case "cde":
		fmt.Println("s is cde")	
	default:
		fmt.Println("unknown")	
	}

//	switch {
//	case s == "abc":
//		fmt.Println("s is abc")
//	case s > "cde":
//		fmt.Println("s is cde")	
//	default:
//		fmt.Println("unknown")	
//	}

}