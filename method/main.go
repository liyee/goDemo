package main

import (
	"fmt"
)

type Circle struct {
	redius float64
}

func (c Circle) area() float64 {
	return c.redius * c.redius * 3.14
}

func main() {
	var c1 Circle
	c1.redius = 10.00
	fmt.Println("面积是", c1.area())

}
