package main

import "fmt"

func search(nums []int, target int) int {
	if len(nums) < 1 {
		return -1
	}
	l := 0
	r := len(nums) - 1
	for l <= r {
		if target < nums[l] && target > nums[r] {
			return -1
		}
		mid := (l + r) / 2
		if nums[l] == target {
			return l
		} else if nums[l] > target {
			if nums[r] == target {
				return r
			}
			if nums[mid] > nums[l] {
				l = mid + 1
			} else {
				if nums[mid] == target {
					return mid
				} else if nums[mid] > target {
					r = mid - 1
				} else {
					l = mid + 1
				}
			}
		} else {
			if nums[mid] < nums[l] {
				r = mid - 1
			} else {
				if nums[mid] == target {
					return mid
				} else if nums[mid] > target {
					r = mid - 1
				} else {
					l = mid + 1
				}
			}
		}
	}
	return -1
}

func main() {
	nums := []int{1}
	fmt.Printf("nums:%+v\n", nums)
	value := search(nums, 1)
	fmt.Printf("value:%d\n", value)
}
