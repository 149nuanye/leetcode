package main

import "fmt"

func lengthOfLIS(nums []int) int {
	maxLength := 0
	if len(nums) == 0 {
		return maxLength
	}
	dp := make([]int, 0, len(nums))
	maxLength = 1
	dp = append(dp, 1)
	for i := 1; i < len(nums); i++ {
		tmp := 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j]+1 > tmp {
					tmp = dp[j] + 1
				}
			}
		}
		if tmp > maxLength {
			maxLength = tmp
		}
		dp = append(dp, tmp)
	}
	return maxLength
}

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	length := lengthOfLIS(nums)
	fmt.Printf("nums:%+v\n", nums)
	fmt.Printf("length:%+v\n", length)
}
