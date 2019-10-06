package main

import (
	"fmt"
)

func main() {
	a := 10
	b := 20

	fmt.Println(a,b)
	swap(&a, &b)
	fmt.Println(a,b)

	//a := 0
	//p := &a  //a의 주소값 c언어랑 비슷한데 주소값 계산은 안됨

	//fmt.Println(&a, *p)
}

func swqp(a, b *int) {
	temp = *a
	*a = *b		
	*b = temp	
	
}