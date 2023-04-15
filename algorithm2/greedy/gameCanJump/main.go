package main

import "fmt"

//游戏跳转
func main() {
	arr := [][]int{
		// {2, 3, 1, 1, 4},
		// {3, 2, 1, 0, 4},
		// {0, 2, 3},
		{2, 3, 1, 2, 4, 2, 3},
	}

	for i := 0; i < len(arr); i++ {
		res := canJump(arr[i])
		fmt.Println(res)

		res2 := jump(arr[i])
		fmt.Println(res2)
	}
}

//力扣地址: https://leetcode.cn/problems/jump-game-ii/
func jump(nums []int) int {
	length := len(nums)
	end := 0
	maxPosition := 0
	steps := 0
	for i := 0; i < length-1; i++ {
		maxPosition = max(maxPosition, i+nums[i])
		if i == end {
			end = maxPosition
			steps++
		}
	}
	return steps
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//力扣地址: https://leetcode.cn/problems/jump-game/
func canJump(nums []int) bool {
	max := 0
	for i := 0; i < len(nums); i++ {
		if max < i {
			return false
		}

		if nums[i]+i > max {
			max = nums[i] + i
		}
	}
	return true
}
