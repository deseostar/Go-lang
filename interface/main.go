package main

import (
	"fmt"
	"reflect"
	"errors"
)

type Runner interface {
	Run()
}

type Gopher struct {
	Name string
	Age int
}

func (g Gopher) Run() {
	fmt.Println("RUN:", g.Name)
}


type Turtle struct {
	Name string
}

func (t Turtle) Run() {
	fmt.Println("SLOW...", t.Name)
}

func (t Turtle) Error() string {
	return t.Name
}

func PrintErr(err error) {
	fmt.Println("ERR:", err.Error())
}


func main() {
	var runner Runner = Gopher {
		Name: "Tom",
		Age: 100,
	}
	fmt.Println(runner) // {Tom 100}
	fmt.Println(reflect.TypeOf(runner))
	runner.Run() // RUN: Tom

	runner = Turtle{
		Name: "TT",
	}
	runner.Run() //SLOW... TT

	// err := errors.New("err")
	// 	error
	// fmt.Println(err)
	
	// t := Turtle{
	// 	Name: "Tom",
	// }
	// printErr(t) // ERR: Tom

	var i interface{}  // 다받을수있지만 사용은 자제하는것이 좋음

	i = bool
	i = 1000
	i = "string"
	i = Turtle{}

	fmt.Println(i)

}
