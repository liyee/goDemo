package main

import "fmt"

type Peple interface {
	Introduction()
}

type Stutent struct {
	Name string
}

func (this *Stutent) Introduction() {
	fmt.Println(this.Name, "is a stutent")
}

func main() {
	peple := &Stutent{"zhangsan"}

	peple.Introduction()
}
