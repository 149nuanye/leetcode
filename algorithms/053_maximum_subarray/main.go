package main

import "fmt"

// 动态规划
// https://leetcode-cn.com/problems/maximum-subarray/solution/zui-da-zi-xu-he-by-leetcode/
// max[i] = Max(max[i-1] + nums[i], nums[i])
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxValue := nums[0]
	tmpMax := nums[0]
	for index := 1; index < len(nums); index++ {
		if tmpMax+nums[index] > nums[index] {
			tmpMax = tmpMax + nums[index]
		} else {
			tmpMax = nums[index]
		}
		if tmpMax > maxValue {
			maxValue = tmpMax
		}
	}
	return maxValue
}

func main() {
	// nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	nums := []int{-2}
	max := maxSubArray(nums)
	fmt.Printf("nums:%+v\nmax:%d\n", nums, max)
}
