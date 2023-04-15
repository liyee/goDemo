package main

import (
	"fmt"
	"time"
)

func say(n int) {
	for i := 0; i < 3; i++ {
		mm := time.Millisecond
		time.Sleep(100 * mm)
		fmt.Println(n)
	}
}

func main() {
	go say(1)
	say(2)
}
