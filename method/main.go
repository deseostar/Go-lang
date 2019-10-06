package main

import (
	"fmt"
)

type Gopher struct {
	Name string
	Age int
}

func (g Gopher) Info() {
	fmt.Printf("[%s] %d\n", g.Name, g.Age)
}

func Info(g, Gopher) {
	fmt.Printf("[%s] %d\n", g.Name, g.Age)
}

func Rename(g *Gopher, newName string) {
	g.Name = newName
}

func (g *Gopher) Rename(newName string) {
	g.Name = newName
}

func main() {
	g := Gopher {
		Name : "Tom",
		Age: 100,
	}

	Info(g) // [Tom] 100  그냥 함수
	g.Info() // [Tom] 100  메서드 리시버를 사용한 고퍼의 메소드
	Gopher.Info(g) // [Tom] 100 고퍼의 메소드를 바로 사용

	Rename(&g, "Pike")

	g.Info() // [Pike] 100

	g. Rename("John")
	g.Info() // [John] 100




}