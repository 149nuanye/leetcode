package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/house-robber/solution/da-jia-jie-she-by-leetcode/
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}
	m := 0
	n := nums[0]
	for index := 1; index <= len(nums)-1; index++ {
		tmp := n
		if m+nums[index] > n {
			n = m + nums[index]
		}
		m = tmp
	}
	if m > n {
		return m
	}
	return n
}

func main() {
	nums := []int{1, 2, 3, 1}
	maxNum := rob(nums)
	fmt.Printf("nums:%+v,maxNum:%d\n", nums, maxNum)
}
