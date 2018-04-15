package main

import "fmt"

type foo struct {
	a uint
	b float32
}

func main() {
	res := foo{a: 2, b: 2.2}

	fmt.Println(res.a, res.b)
}
