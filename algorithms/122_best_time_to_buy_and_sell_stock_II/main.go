package main

import "fmt"

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/solution/mai-mai-gu-piao-de-zui-jia-shi-ji-ii-by-leetcode/
func maxProfit(prices []int) int {
	profit := 0
	for index := 1; index < len(prices); index++ {
		if prices[index]-prices[index-1] > 0 {
			profit += prices[index] - prices[index-1]
		}
	}
	return profit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	profit := maxProfit(prices)
	fmt.Printf("prices:%+v,profit:%d\n", prices, profit)
}
