package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func say(s string) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 200)
	}
}

func main() {
	wg.Add(1)
	go say("hi")
	wg.Add(1)
	go say("there")
	wg.Wait()
}
