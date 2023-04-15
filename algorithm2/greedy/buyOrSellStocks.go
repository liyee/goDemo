package main

import "fmt"

/*
股票交易：通过局部最优得到全局最优
*/
func main2() {
	arr := [][]int{
		{7, 1, 5, 3, 6, 4},
		{1, 2, 3, 4, 5},
		{7, 6, 4, 3, 1},
	}

	for i := 0; i < len(arr); i++ {
		num := maxProfit(arr[i])
		fmt.Println(num)
	}
}

func maxProfit(prices []int) int {
	profit := 0

	for i := 1; i < len(prices); i++ {
		diff := prices[i] - prices[i-1]
		if diff > 0 {
			profit += diff
		}
	}

	return profit
}

//针对力库，股票一次交易：https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/
func maxProfit2(prices []int) int {
	minprice := prices[0]
	maxprofit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < minprice {
			minprice = prices[i]
		} else if prices[i]-minprice > maxprofit {
			maxprofit = prices[i] - minprice
		}
	}
	return maxprofit
}
