package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	values := make([][]int, 0, 10)
	sort.Ints(nums)
	for index := 0; index < len(nums); index++ {
		if nums[index] > 0 {
			break
		}
		if index > 0 && nums[index] == nums[index-1] {
			continue
		}
		l := index + 1
		r := len(nums) - 1
		for l < r {
			sum := nums[index] + nums[l] + nums[r]
			if sum == 0 {
				values = append(values, []int{nums[index], nums[l], nums[r]})
				for l < r {
					if nums[l+1] == nums[l] {
						l++
						continue
					}
					break
				}
				for l < r {
					if nums[r] == nums[r-1] {
						r--
						continue
					}
					break
				}
				l++
				r--
			} else if sum > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return values
}

func main() {
	nums := []int{4, -1, -1, 0, 1, 2}
	values := threeSum(nums)
	for index := range values {
		fmt.Printf("value:%+v\n", values[index])
	}
}
