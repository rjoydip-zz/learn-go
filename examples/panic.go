package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i == 2 {
			panic("I am panic in 2")
		}
	}
}
