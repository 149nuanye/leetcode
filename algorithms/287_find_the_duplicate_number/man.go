package main

import "fmt"

func findDuplicate(nums []int) int {
	fast := 0
	slow := 0
	for fast < len(nums) && nums[fast] < len(nums) {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	fast = 0
	for fast < len(nums) {
		if nums[fast] == nums[slow] {
			return nums[fast]
		}
		fast = nums[fast]
		slow = nums[slow]
	}
	return 0
}

func main() {
	nums := []int{1, 3, 4, 2, 2}
	dupNum := findDuplicate(nums)
	fmt.Printf("nums:%+v,dupNum:%d\n", nums, dupNum)
}
