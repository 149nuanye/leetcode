package main

import "fmt"

// https://leetcode.com/problems/two-sum/
func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}
	dataMap := make(map[int]int)
	result := make([]int, 0, 2)
	for index := range nums {
		other := target - nums[index]
		if _, ok := dataMap[other]; ok {
			result = append(result, dataMap[other])
			result = append(result, index)
			break
		}
		dataMap[nums[index]] = index
	}
	return result
}

func main() {
	nums := []int{3, 7, 11, 15, 2}
	result := twoSum(nums, 9)
	fmt.Printf("result:%+v\n", result)
}
