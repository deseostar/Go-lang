package main

import (
	"fmt"
)

func main() {
	// a:= []int{1,2,3,4,}
	// fmt.Println(a)
	// fmt.Println(len(a), cap(a))

	// a = append(a, 5)
	// fmt.Println(a)
	// fmt.Println(len(a), cap(a))  // 기존 array 용량의 두배가 됨 그래서 8이 됨

	a := make([]int, 4, 5)  // 추가적인 할당과 복사를 막을수 있다. 이래서 make통해 직접 넣음
	fmt.Println(a)  //[0 0 0 0]
	fmt.Println(len(a), cap(a))  //4 5
}