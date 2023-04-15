package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type man struct {
	Name string `json:name`
	Age  int    `json:age`
	Sex  string `json:sex`
}

func (this man) PrintStr() string {
	return this.Name + "+" + strconv.Itoa(this.Age) + "+" + this.Sex
}

func main() {
	manList := man{"a", 18, "c"}
	val := reflect.ValueOf(manList)
	v := val.Method(0).Call(nil)
	fmt.Println(v)
	fmt.Println(val.MethodByName("PrintStr").Call(nil))

	//获取字段名称
	// manType := reflect.TypeOf(manList)
	// for i := 0; i < 3; i++ {
	// 	val := manType.Field(i)
	// 	fmt.Printf("%v, %v\n", val.Name, val.Tag)
	// }
}
