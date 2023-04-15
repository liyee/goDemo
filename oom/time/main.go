package main

import (
	"fmt"
	"time"
)

func main() {
	// fmt.Println("Hello World!")
	// mutexText()

	//1.无限定时循环
	/*
	  tiker1 := time.NewTicker(1 * time.Second)
	  	defer tiker1.Stop()
	  	for range tiker1.C {
	  		fmt.Println("ok")
	  	}
	*/

	// 2.限定循环次数执行
	/*
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()
		go func() {
			for i := 0; i <= 5; i++ {
				fmt.Println(<-ticker.C)
				if i == 5 {
					ticker.Stop()
				}
			}

		}()
		for {
		}
	*/

	/*
	   timer1 := time.NewTimer(2 * time.Second)
	   	t1 := time.Now()
	   	fmt.Println(t1)

	   	t2 := <-timer1.C
	   	fmt.Println(t2)

	   	timer2 := time.NewTimer(time.Second)
	   	for {
	   		<-timer2.C
	   		fmt.Println("时间到，当前for循环不再执行")
	   	}

	*/

	timer := time.NewTicker(time.Duration(2) * time.Second)
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			fmt.Println("ok")
		default:
			fmt.Println("default")
			time.Sleep(time.Duration(1) * time.Second)
		}
	}
}

// func mutexText() {
// 	mutex := sync.Mutex{}
// 	for i := 0; i < 10; i++ {
// 		go func() {
// 			mutex.Lock()
// 			fmt.Printf("%d goroutine is run", i)
// 			time.Sleep(5 * time.Microsecond)
// 		}()

// 	}
// }
