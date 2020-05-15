package main

import "fmt"

func maxProfit(prices []int) int {
	maxP := 0
	if len(prices) == 0 {
		return maxP
	}
	minValue := prices[0]
	for index := 1; index < len(prices); index++ {
		if prices[index] < minValue {
			minValue = prices[index]
			continue
		}
		if prices[index]-minValue > maxP {
			maxP = prices[index] - minValue
		}
	}
	return maxP
}

func main() {
	nums := []int{7, 1, 5, 3, 6, 4}
	maxP := maxProfit(nums)
	fmt.Printf("nums:%+v,maxProfit:%+v\n", nums, maxP)
}
