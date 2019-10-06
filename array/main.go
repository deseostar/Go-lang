package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var arr [10]int

	// fmt.Println(arr)
	// 결과값 [0 0 0 0 0 0 0 0 0 0]

	arr := [5]int{1,2,3,4,5}
	fmt.Println(reflect.TypeOf(arr))
	arr2 := [5]int{1,2,3,4,5}

	fmt.Println(arr == arr2) // true

	fmt.Println(arr)
	fmt.Println(arr[0])
	fmt.Println(arr[2])
	fmt.Println(arr[4])
	// 결과값[1 2 3 4 5]
	// 1
	// 3
	// 5

	Change(&arr, 1, 1000)
	fmt.Println(arr)
	// [1 2 3 4 5]

	slice := arr[0:3]
	fmt.Println(slice)
	fmt.Println(reflect.TypeOf(slice))

}

func Change(arr *[5]int, idx int, val int) {
	arr[idx] = val
}