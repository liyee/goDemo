package main

import (
	"fmt"
	"math/rand"
	"time"
)

// "crypto/rand"

func main() {
	// num, _ := rand.Int(rand.Reader, big.NewInt(5))
	// fmt.Println(num)

	for i := 0; i < 3; i++ {
		rand.Seed(time.Now().UnixNano())
		num2 := rand.Intn(3)
		fmt.Println(num2)
	}

}
