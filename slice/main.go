package main

import "fmt"

func main() {
	//两个切片是否相等
	// s1 := []int{1, 2, 3, 4}
	// s2 := []int{1, 2, 3, 4, 5}
	// fmt.Println(reflect.DeepEqual(s1, s2))

	nums := []int{1, 2}
	fmt.Println(nums, len(nums), cap(nums))

	nums = append(nums, 3, 4, 5)
	fmt.Println(nums, len(nums), cap(nums))

}
