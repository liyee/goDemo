package main

// "fmt"

func fibonacci(n int, ch chan int) {
	x, y := 0, 1

	for i := 0; i < n; i++ {
		ch <- x

		x, y = y, x+y
	}
	close(ch)
}

// func main() {
// 	ch := make(chan int, 5)
// 	go fibonacci(cap(ch), ch)

// 	for v := range ch {
// 		fmt.Println(v)
// 	}
// }
