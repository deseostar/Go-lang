package main

import (
	"fmt"
)

func main() {
	var cnt int

	for i := 0; i < 10; i++ {
		cnt++
	}


	for cnt < 15 {  //while가 없는대신 사용
		cnt++
	}

	//for {  //while가 없는대신 사용
	//	cnt++
	//}

	// for true {
	//	break
	// }

	fmt.Println(cnt)
}