package main

import (
	"time"
)

func f1(ch chan struct{}) {
	time.Sleep(1 * time.Second)
	ch <- struct{}{}
}

func f2(ch chan struct{}) {
	time.Sleep(3 * time.Second)
	ch <- struct{}{}
}

// func main() {
// 	ch1, ch2 := make(chan struct{}), make(chan struct{})
// 	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
// 	go func() {
// 		go f1(ch1)
// 		select {
// 		case <-ch1:
// 			fmt.Println("f1 is done")
// 		case <-ctx.Done():
// 			fmt.Println("f1 timeout")
// 			break
// 		}
// 	}()

// 	go func() {
// 		go f2(ch2)
// 		select {
// 		case <-ch2:
// 			fmt.Println("f2 is done")
// 		case <-ctx.Done():
// 			fmt.Println("f2 timeout")
// 			break
// 		}
// 	}()

// 	ctx.Done()
// 	time.Sleep(5 * time.Second)
// }
