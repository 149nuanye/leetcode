package main

import "fmt"

func canJump(nums []int) bool {
	if len(nums) < 1 {
		return true
	}
	longJump := nums[0]
	for index := 1; index < len(nums); index++ {
		if longJump >= len(nums)-1 {
			return true
		}
		if longJump < index {
			return false
		}
		if index+nums[index] > longJump {
			longJump = index + nums[index]
		}
	}
	return true
}

func main() {
	nums := []int{2, 3, 1, 1, 4}
	jump := canJump(nums)
	fmt.Printf("nums:%+v,jump:%+v\n", nums, jump)
}
