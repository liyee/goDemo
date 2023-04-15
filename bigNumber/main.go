package main

import (
	"fmt"
	// "log"
	"math/big"
)

func main() {
	i := new(big.Int)
	i.SetString("40000000000000000000000000000000000", 10)
	i.Mul(i, big.NewInt(4))
	fmt.Println(i)

	// var n int64 = 2 << 70
	// a := big.NewInt(n)
	// fmt.Println(a)

	// i := new(big.Int)
	// _, err := fmt.Sscan("18446744073709551617", i)
	// if err != nil {
	// 	log.Println("error scanning value:", err)
	// } else {
	// 	fmt.Println(i)
	// }
}
