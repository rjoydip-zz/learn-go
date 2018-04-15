package main

import "fmt"

type foo struct {
	a uint16 // min 0 max 65535
	b int16  // -32k - + 32k
	c float32
}

func main() {
	res := foo{a: 2, b: 2, c: 2.2}

	fmt.Println(res.a, res.b)
}
