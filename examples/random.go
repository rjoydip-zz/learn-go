package main

import (
	"fmt"
	"math/rand"
	"time"
)

func _rand() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	rn := r1.Intn(100)
	fmt.Println("Random number from (1-100): ", rn)
}

func main() {
	_rand()
}
