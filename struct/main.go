package main

import (
	"fmt"
)

type Gopher struct {
	Name string
	Age int
	Phone Phone
}

// type Animal struct {  // 상속이랑 비슷
// 	Name string
// 	Age int
// }

// type Phone struct {
// 	Animal           // 인베딩 상속이랑 비슷
// 	Model string
// }

func main() {
	// // g := Gopher{"Name", 19}
	
	//g := Gopher{
	//	Name: "Name",
	//	Age: 19,
	//
	//	// fmt.Println(g == Gopher{})
	//}
	//fmt.Println(g)

	p := Phone{
		Model: "iPhone",
	}
	g := Gopher{
		Name: "Pike",
		Age : 50,
		Phone: p,
	}
	fmt.Println(g)
	fmt.Println(g.Name, g.Age, g.Phone)
}