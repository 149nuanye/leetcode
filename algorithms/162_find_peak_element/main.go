package main

import "fmt"

// https://leetcode-cn.com/problems/find-peak-element/solution/xun-zhao-feng-zhi-by-leetcode/
func findPeakElement(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return i
		}
	}
	return len(nums) - 1
}

func main() {
	nums := []int{1, 2, 1, 3, 5, 6, 4}
	peekIndex := findPeakElement(nums)
	fmt.Printf("nums:%+v, peekIndex:%d\n", nums, peekIndex)
}
