package main

import "fmt"

func cleanup() {
	if r := recover(); r != nil {
		fmt.Println("Recovered in cleanup: ", r)
	}
}

func main() {
	defer cleanup()
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i == 2 {
			panic("I am panic in 2")
		}
	}
}
