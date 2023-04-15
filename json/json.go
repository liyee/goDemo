package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Auth  string  `json:"auth"`
	Price float32 `json:"price"`
}

func main() {
	book1 := Book{
		"php",
		32.00,
	}

	jsonstr, err := json.Marshal(book1)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	fmt.Printf("book1: %s\n", jsonstr)

	myBook := Book{}
	err = json.Unmarshal(jsonstr, &myBook)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	fmt.Printf("%v", myBook)
}
