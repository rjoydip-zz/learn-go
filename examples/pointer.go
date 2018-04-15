package main

import (
	"fmt"
)

func main() {
	x := 12
	a := &x // memory address

	fmt.Println(*a, a)
	*a = 5
	fmt.Println(*a, a, x)
}
