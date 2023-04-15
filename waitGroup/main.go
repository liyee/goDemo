package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var taskCount, count = 100, 10
	c := make(chan struct{}, count)
	defer close(c)

	for i := 0; i < taskCount; i++ {
		wg.Add(1)
		c <- struct{}{}
		go func(j int) {
			defer wg.Done()
			fmt.Println(j, time.Now().UnixNano())
			<-c
		}(i)

	}

	wg.Wait()

}
