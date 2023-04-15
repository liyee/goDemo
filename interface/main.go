package main

import (
	"fmt"
)

type Phone interface {
	call()
}

type Nokia struct {
}

func (ph Nokia) call() {
	fmt.Println("nokia")
}

type Stu struct {
	name string
}

type StuInt interface{}

func main() {
	var phone Phone
	phone = new(Nokia)
	phone.call()

	//比较两个interface是否相等
	fmt.Println("\n比较两个interface是否相等")
	var str1, str2 StuInt = &Stu{"zhang"}, &Stu{"zhang"}
	var str3, str4 StuInt = Stu{"zhang"}, Stu{"zhang"}

	fmt.Println(str1 == str2)
	fmt.Println(str3 == str4)

	var p *int = nil
	var i interface{} = nil
	if p == i {
		fmt.Println("Equal")
	} else {
		fmt.Printf("%T, %T: ", p, i)
		fmt.Println("Not Equal")
	}
}
