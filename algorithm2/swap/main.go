package main

import (
	"fmt"
)

func main() {
	test(10009)
}

func test(N int) {
	enable_print := N % 10
	for N > 0 {
		if enable_print == 0 && N%10 != 0 {
			enable_print = 1
		}

		if enable_print < 10 {
			fmt.Print(N % 10)
		}
		N = N / 10
	}
}
