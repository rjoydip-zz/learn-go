package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 200)
	}
	wg.Done()
}

func main() {
	wg.Add(1)
	go say("hi")
	wg.Add(1)
	go say("there")
	wg.Wait()

	wg.Add(1)
	go say("Hello")
	wg.Wait()
}
