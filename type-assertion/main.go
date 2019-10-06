package main

import (
	"fmt"
)

type Sounder interface {
	Sound() string
}

type Pig struct {
	Name string
}

func (p Pig) Sound() string {
	return "꿀"
}

type Duck struct {
	Name string
}

func (d Duck) Sound() string {
	return "꽥"
}

func PlaySound(s Sounder) {
	// val, ok := s.(Pig)
	// if !ok {
	// 	fmt.Println("It's not a pig.")
	// }
	// fmt.Println(val.Sound())
	switch val := s.(type) {
	case Pig:
		fmt.Println("돼지:", val.Sound())
	case Duck:
		fmt.Println("오리:", val.Sound())
	default:
		fmt.Println("UNKNOWN:", s.Sound())
	}

}

type Gopher string
func (g Gopher) Sound() string {
	return string(g)
}

func main() {

	p := Pig{}
	d := Duck{}
	g := Gopher("Tom")

	PlaySound(p) // 꿀
	PlaySound(d) // 꽥
	PlaySound(g)
}

// 돼지: 꿀
// 오리: 꽥
// UNKNOWN: Tom