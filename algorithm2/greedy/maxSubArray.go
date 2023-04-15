package main

import (
	"fmt"
)

/*最大子序和*/
func main() {
	arr := [][]int{
		{-2, 1, -3, 4, -1, 2, 1, -5, 4},
		{1},
		{-1},
		{5, 4, -1, 7, 8},
		{-2, 1},
		{-1, 1, 2, 1},
	}

	for i := 0; i < len(arr); i++ {
		num := maxSubArray(arr[i])
		fmt.Println(num)
	}
}

func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}

		if nums[i] > max {
			max = nums[i]
		}
	}

	return max
}
