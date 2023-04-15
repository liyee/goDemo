package main

import (
	"fmt"
)

// func main() {
// 	cat, fish, dog := make(chan struct{}), make(chan struct{}), make(chan struct{})

// 	for i := 0; i < 100; i++ {
// 		go Cat(cat, fish, i)
// 		go Fish(fish, dog)
// 		go Dog(dog, cat)
// 	}

// 	cat <- struct{}{}
// 	time.Sleep(5 * time.Second)
// }

func Cat(cat, fish chan struct{}, i int) {
	<-cat
	fmt.Print(i, ": cat, ")
	fish <- struct{}{}
}

func Fish(fish, dog chan struct{}) {
	<-fish
	fmt.Print("fish, ")
	dog <- struct{}{}
}

func Dog(dog, cat chan struct{}) {
	<-dog
	fmt.Println("dog")
	cat <- struct{}{}
}
