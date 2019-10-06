package main

import (
	"fmt"
	//"io"
)

func main() {
	fmt.Println("main 1")

	defer fmt.Println("defer 1")

	fmt.Println("main 2")

	defer fmt.Println("defer 2")

	fmt.Println("main 3")

	//main 1  
	//main 2
	//main 3
	//defer 2
	//defer 1

	//r := io.ReadCloser(nil) 결과값 순서가 특이하네 (쓰임새: 리턴이 여러곳에서 일어날때 
	//r.Close()  				그 함수를 결국 닫아줄때 위에서 닫아줄수 있다. 보기 좋은듯)
	// defer func() {   또, 결과가 나오기전에 값을 변경 해줄 수도 있음
	//	res = 100
	//}


}