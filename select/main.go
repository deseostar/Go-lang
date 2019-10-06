package main

import (
	"fmt"
)

type Sender struct {
	val string
}

func (s Sender) Send(c chan<- string) { //이렇게 하면 값을 보낼수만 있음
	c <- s.val
}

func main() {
	s1 := Sender{
		val: "AAA",
	}
	s2 := Sender{
		val: "BBB",
	}

	c1 := make(chan string)
	c2 := make(chan string)

	go s1.Send(c1)
	go s2.Send(c2)

	for {
		select {
		case val := <- c1:
			fmt.Println("C1:", val)
		case val := <- c2:
			fmt.Println("C2:", val)
		default :
		fmt.Println("default") // deadlock 블락 되면 계속 찍힘
		}
	}
}