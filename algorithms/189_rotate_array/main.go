package main

import (
	"fmt"
)

func rotateSlow(nums []int, k int) {
	length := len(nums)
	if length == 0 {
		return
	}
	k = k % length
	if k == 0 {
		return
	}
	for index := 0; index < k; index++ {
		last := nums[length-1]
		for i := length - 1; i > 0; i-- {
			nums[i] = nums[i-1]
		}
		nums[0] = last
	}
}

func rotateVector(nums []int) {
	l := 0
	r := len(nums) - 1
	for l < r {
		nums[r], nums[l] = nums[l], nums[r]
		l++
		r--
	}
}

// https://leetcode-cn.com/problems/rotate-array/solution/man-hua-san-ci-xuan-zhuan-de-fang-fa-shi-ru-he-x-2/
func rotate(nums []int, k int) {
	length := len(nums)
	if length == 0 {
		return
	}
	k = k % length
	if k == 0 {
		return
	}
	rotateVector(nums)
	rotateVector(nums[:k])
	rotateVector(nums[k:])
}
func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	fmt.Printf("k:%d,nums:%+v\n", k, nums)
	rotate(nums, k)
	fmt.Printf("nums:%+v\n", nums)
}
