package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	fmt.Println("Hello Go")
	msgList := make(chan int, 100)

	go func() {
		for {
			select {
			case <-msgList:
			default:
			}
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	s := <-c

	fmt.Println("main exit.get signal:", s)
}
