package main

import (
	"fmt"
)

func changeValue(myarr []int) {
	capNum := 0
	for i := 11; i <= 1000; i++ {
		myarr = append(myarr, i)
		if capNum != cap(myarr) {
			capNum = cap(myarr)
			fmt.Println("new: ", len(myarr), cap(myarr))
		}
	}

}
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	changeValue(arr)
	fmt.Println("====")
	fmt.Println("old: ", len(arr), cap(arr))
}
