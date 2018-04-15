package main

import "fmt"

func foo() {
	defer fmt.Println("Next after second")
	defer fmt.Println("Next after first")
	fmt.Println("Last line come first")
}

func main() {
	foo()
}
