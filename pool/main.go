package main

import (
	"fmt"
	"sync"
)

func main() {
	count := 100
	ch := make(chan struct{}, 10)
	wg := sync.WaitGroup{}
	defer close(ch)

	for i := 0; i < count; i++ {
		wg.Add(1)
		ch <- struct{}{}
		go func(j int) {
			defer wg.Done()
			fmt.Println(j)
			<-ch
		}(i)
	}

	wg.Wait()

}
