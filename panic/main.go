package main

import (
	"fmt"
	"errors"
)

func div(a,b int) (res int) {
	defer func() {
		if err := recover(); err != nil {   //recober 약간 트라이 캐치랑 비슷
			fmt.Printls(err)
			res = -1
		}
	}()

	res = a / b

	return
}

func main() {
	fmt.Println(div(1,2))
	fmt.Println(div(10, 0))
}