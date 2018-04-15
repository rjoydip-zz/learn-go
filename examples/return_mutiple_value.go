package main

import (
	"fmt"
)

func returnMultipleValue(x, y string) (string, string) {
	return x, y
}

func main() {
	num1, num2 := "Hello", "world"

	fmt.Println(returnMultipleValue(num1, num2))
}
