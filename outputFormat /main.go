package main

import (
	"fmt"
)

type peple struct {
	name string
	age  int
	sex  uint8
}

func main() {
	p := peple{"zhangsan", 18, 1}
	fmt.Printf("%v\n", p)
	fmt.Printf("%+v\n", p)
	fmt.Printf("%#v\n", p)
}
