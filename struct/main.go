package main

import (
	"fmt"
)

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	// fmt.Println(Books{"golang", "ljq", "go 语言教程", 1})
	chinese := Books{"golang", "ljq", "go 语言教程", 1}
	fmt.Println(chinese)

	printBook(&chinese)

	// fmt.Println(chinese, chinese.title)
}

func printBook(book *Books) {
	fmt.Println(book.title, book.author, book.subject, book.book_id)
}
