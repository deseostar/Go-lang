package main

import (
	//f "fmt"
	//. "strings"  네임스페이스  .
	//_ "strings"  import만 하고 사용안할때 _
	"fmt"
	"GoEasy/package/math"
)

func main() {
	//fmt.Println("HI")
	//f.Println(strings.Repeat("Hi", 5))
	fmt.Println(math.Add(1, 2)) //3   // Add의 맨앞 대문자로 해야 외부에서 사용가능
	fmt.Println(math.N)  //1
}