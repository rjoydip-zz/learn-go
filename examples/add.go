package main

import (
	"fmt"
)

func add(x, y float64) float64 {
	return x + y
}

func main() {
	var num1, num2 float64 = 2.3, 3.2

	fmt.Println("Addition of (2.3 + 3.2)", add(num1, num2))
}
