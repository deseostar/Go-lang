package main

import (
	"fmt"
)

func main() {  // 대표 3가지 array slice map
	a := [5]int{1,2,3,4,5}   
	s := []int{100, 200, 300, 400, 500}
	m := map[string]int{
		"one" : 1,
		"two" : 2,
		"three" : 3,
		"four": 4,
		"five": 5,
	}

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}


	for idx := range a {  // range가 더 사용이 편리하며 map의 경우에도 사용가능
		fmt.Println(a[idx])		
	}
	for idx, val := range s {
		fmt.Println(idx, val)
	} 
	for key, val := range m {
		fmt.Println(key, val)
	}

	// for key := range m {  // 맵은 순서가 없음
	// 	delete(m, key)
	// }
	// fmt.Println(m)

// 결과값
// 1
// 2
// 3
// 4
// 5
// 100
// 200
// 300
// 400
// 500
// 1
// 2
// 3
// 4
// 5
// 0 100
// 1 200
// 2 300
// 3 400
// 4 500
// three 3
// four 4
// five 5
// one 1
// two 2
}