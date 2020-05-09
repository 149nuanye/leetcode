package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	value := nums[0]
	lindex := 0
	for rindex := 1; rindex < len(nums); rindex++ {
		if value == nums[rindex] {
			continue
		} else {
			value = nums[rindex]
			lindex++
			nums[lindex] = nums[rindex]
		}
	}
	nums = nums[:lindex+1]
	return lindex + 1
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Printf("nums:%+v\n", nums)
	newLength := removeDuplicates(nums)
	fmt.Printf("newLength:%d\n", newLength)
	for index := 0; index < newLength; index++ {
		fmt.Printf("num:%+v\n", nums[index])
	}
}
