package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{7, 2, 2, 3, 4, 5, 6}

	f1(arr, 10)
}

var (
	res  [][]int
	path []int
)

func f1(arr []int, target int) {
	sort.Ints(arr)
	dfs(arr, 10, 0)
	fmt.Println(res)
}

func dfs(arr []int, target int, start int) {
	if target == 0 {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		return
	}

	for i := start; i < len(arr); i++ {
		if i != start && arr[i-1] == arr[i] {
			break
		}

		path = append(path, arr[i])
		dfs(arr, target-arr[i], i+1)
		path = path[:len(path)-1]
	}
}
