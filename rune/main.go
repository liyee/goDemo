package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "hello world! 您好，世界！"

	arr := []rune(str)

	fmt.Println(string(arr), len(arr), utf8.RuneCountInString(str), len(str), str[10:30], string(arr[:19]))
}
