package main

import "fmt"

// https://leetcode-cn.com/problems/first-missing-positive/solution/que-shi-de-di-yi-ge-zheng-shu-by-leetcode/
func firstMissingPositive(nums []int) int {
	contains := 0
	for index := range nums {
		if nums[index] == 1 {
			contains++
		}
	}
	if contains == 0 {
		return 1
	}
	if len(nums) == 1 {
		return 2
	}

	n := len(nums)
	for index := range nums {
		if nums[index] <= 0 || nums[index] > n {
			nums[index] = 1
		}
	}

	for index := range nums {
		tmp := nums[index]
		if tmp < 0 {
			tmp = -tmp
		}
		if tmp == n {
			if nums[0] > 0 {
				nums[0] = -nums[0]
			}
		} else {
			if nums[tmp] > 0 {
				nums[tmp] = -nums[tmp]
			}
		}
	}

	for index := range nums {
		if index == 0 {
			continue
		}
		if nums[index] > 0 {
			return index
		}
	}
	if nums[0] > 0 {
		return len(nums)
	}
	return len(nums) + 1
}

func firstMissingPositiveHash(nums []int) int {
	values := make(map[int]bool)
	for index := range nums {
		if _, ok := values[nums[index]]; !ok {
			values[nums[index]] = true
		}
	}
	for index := 1; index <= len(nums); index++ {
		if _, ok := values[index]; !ok {
			return index
		}
	}
	return len(nums) + 1
}

func main() {
	values := []int{3, 4, 2, 1}
	fmt.Printf("values:%+v\n", values)
	fmt.Printf("dstValue:%d\n", firstMissingPositive(values))
}
