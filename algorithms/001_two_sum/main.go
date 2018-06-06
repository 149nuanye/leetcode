package main

import (
	"fmt"
)

// desc https://leetcode.com/problems/two-sum/description/

func twoSum(nums []int, target int) []int {
	result := make([]int, 0, 2)
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = append(result, i, j)
				return result
			}
		}
	}
	return result
}

func twoSum2(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i := range nums {
		numMap[nums[i]] = i
	}
	result := make([]int, 0, 2)
	for i := range nums {
		key := target - nums[i]
		if _, ok := numMap[key]; ok && numMap[key] != i {
			result = append(result, i, numMap[key])
			return result
		}
	}
	return result
}

func main() {
	nums := []int{3, 2, 4}
	target := 6
	res := twoSum2(nums, target)
	fmt.Printf("ret result:%v\n", res)
}
