package main

import "fmt"

func main() {
	ch1, ch2 := make(chan string, 1), make(chan string, 1)
	ch1 <- "a"
	ch2 <- "b"

	prioritySlect(ch1, ch2)

}

func prioritySlect(ch1, ch2 <-chan string) {
	for {
		select {
		case val := <-ch1:
			fmt.Println(val)
		case val2 := <-ch2:
		priority:
			for {
				select {
				case val1 := <-ch1:
					fmt.Println(val1)

				default:
					break priority
				}
			}
			fmt.Println(val2)
		}
	}

}
