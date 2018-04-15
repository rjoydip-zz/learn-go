package main

import "fmt"

type foo struct {
	a uint
	b float32
}

func (o foo) add() float32 {
	return float32(o.a) * o.b
}

func main() {
	res := foo{a: 2, b: 2.2}

	fmt.Println(res.add())
}
