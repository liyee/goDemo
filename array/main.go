package main

import (
	"fmt"
)

func changeValue(slice []int) {
	for i := 6; i < 100; i++ {
		slice = append(slice, i)
	}
	slice[0] = 100

	fmt.Printf("%T\n", slice)
}

//1,1,2,3,5,8,13,21,34
func fibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	}

	var a, b int = 1, 1
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}

	return a + b
}

func main() {
	// slice := []int{1, 2, 3, 4, 5}
	// fmt.Printf("%T\n", slice)
	// changeValue(slice)
	// fmt.Println(slice)
	val := fibonacci(7)

	fmt.Println(val)
}
