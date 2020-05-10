package main

import (
	"fmt"
)

func moveZeroes(nums []int) {
	zindex := 0
	for index := 0; index < len(nums); index++ {
		if nums[index] != 0 {
			nums[zindex], nums[index] = nums[index], nums[zindex]
			zindex++
		}
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Printf("nums:%+v\n", nums)
}
