package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	map1 := map[int]string{1: "go", 2: "php", 3: "java", 4: "c++"}
	var lock sync.RWMutex

	for i := 0; i < 500; i++ {
		go func() {
			defer lock.RUnlock()
			fmt.Println(map1[2])
			lock.RLock()
		}()
	}

	end := time.Now()
	fmt.Println(end.Sub(start))
	time.Sleep(2 * time.Second)
}
