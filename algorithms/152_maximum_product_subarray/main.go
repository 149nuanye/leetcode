package main

import "fmt"

// https://leetcode-cn.com/problems/maximum-product-subarray/solution/dpfang-fa-xiang-jie-by-yang-cong-12/
func maxProduct(nums []int) int {
	maxValue := 0
	if len(nums) == 0 {
		return maxValue
	}
	tmpMax := nums[0]
	maxValue = tmpMax
	tmpMin := nums[0]
	for i := 1; i < len(nums); i++ {
		tmp := tmpMax

		if nums[i] > tmpMax*nums[i] {
			tmpMax = nums[i]
		} else {
			tmpMax = tmpMax * nums[i]
		}
		if tmpMax < tmpMin*nums[i] {
			tmpMax = tmpMin * nums[i]
		}

		if nums[i] > tmpMin*nums[i] {
			tmpMin = tmpMin * nums[i]
		} else {
			tmpMin = nums[i]
		}
		if tmpMin > tmp*nums[i] {
			tmpMin = tmp * nums[i]
		}
		// fmt.Printf("value:%d, tmpMax:%d, tmpMin:%d, maxValue:%d\n", nums[i], tmpMax, tmpMin, maxValue)
		if tmpMax > maxValue {
			maxValue = tmpMax
		}
	}
	return maxValue
}

func main() {
	nums := []int{2, -1, 3, -1}
	maxValue := maxProduct(nums)
	fmt.Printf("nums:%+v,maxValue:%d\n", nums, maxValue)
}
