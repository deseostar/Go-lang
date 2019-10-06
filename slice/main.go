package main

import (
	"fmt"
	_ "reflect"   // 안쓰는데 넣어두면 에러남(밑에꺼 주석처리하면서 놓쳤더니 이모양)
)

func main() {
	// s := []int{1,2,3,4,5}

	// fmt.Println(s)
	// fmt.Println(reflect.TypeOf(s))
	// // [1 2 3 4 5]
	// // []int 

	// var s []int
	// fmt.Println(s == nil) //true
	// fmt.Println(reflect.TypeOf(s))

	// s := make([]int, 5)
	// fmt.Println(s == nil) //false
	// fmt.Println(s) //[0 0 0 0 0]
	// fmt.Println(reflect.TypeOf(s))  //[]int

	a := []int{1, 200, 3000, 4000}
	b := a[0: 3]

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(len(a), cap(a))
	fmt.Println(len(b), cap(b))

	// if len(a) == 0 {

	// }  이렇게 쓰면 슬라이스가 비워져 있든 nil이 들어가 있든 둘다 잡아냄

	b[1] = 5000  // 이러면 a랑 b랑 둘다 바뀜 같은 슬라이스를 공유, 조심해서 써야함
				// 그래서 슬라이스끼리는 비교못함, FOR를 써서 하나씩 꺼내 비교해야함??

	fmt.Println(a)
	fmt.Println(b)
}

