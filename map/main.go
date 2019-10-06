package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"A": 1,
		"B": 3000,
		"C": 50000,
		"E": 0,
	}
	m["D"] = 600000

	fmt.Println(m) // map[]
	val := m["D"]
	fmt.Println(val)
	
	
	val = m["E"] //정의되지 않은 값
	fmt.Println(val) // 0 이 나옴 그래서 값이 있는지 아님 0이 들어있는지 확인필요

	// val, ok := m["E"]
	// fmt.Println(val, ok) // 그래서 이거랑 밑에 랑 씀

	if _, ok := m["E"]; ok {
		fmt.Println("E!!")
	}

	delete(m, "E")
	if _, ok := m["E"]; ok {
		fmt.Println("EEE")
	}
}