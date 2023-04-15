package main

import (
	"fmt"
	"strings"
)

/*
459. 重复的子字符串: https://leetcode.cn/problems/repeated-substring-pattern/
*/
func main() {
	arr := []string{
		"abab",
		"aba",
		"abcabcabcabc",
	}

	for i := 0; i < len(arr); i++ {
		res := repeatedSubstringPattern(arr[i])
		fmt.Println(res)
	}
}

func repeatedSubstringPattern(s string) bool {
	fmt.Println(s[1:] + s[:len(s)-1])

	return strings.Contains(s[1:]+s[:len(s)-1], s)

	l := len(s)

	for step := 0; step < l/2; {
	LOOP:
		step++

		if l%step != 0 {
			goto LOOP
		}

		for i := 0; i < l-step; i++ {
			if s[i] != s[i+step] {
				goto LOOP
			}

			if i+step == l-1 {
				return true
			}
		}
	}

	return false
}
