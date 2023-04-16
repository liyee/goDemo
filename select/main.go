package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	result := make(chan int)

	go func() {
		interval := time.Duration(rand.Int31n(10)) * time.Second
		fmt.Println("sleep:", interval)
		time.Sleep(interval)
		result <- int(interval)
	}()

	select {
	case <-result:
		fmt.Println("ok")
	case <-time.After(5 * time.Second):
		fmt.Println("timeout")
	}

}
