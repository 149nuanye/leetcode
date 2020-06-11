package main

import "fmt"

// https://leetcode-cn.com/problems/kth-largest-element-in-an-array/solution/shu-zu-zhong-de-di-kge-zui-da-yuan-su-by-leetcode/
func quickselect(left, right, kSmallest int, nums []int) int {
	fmt.Printf("nums:%+v\n", nums)
	tag := nums[left]
	lindex := left
	rindex := right
	for lindex < rindex {
		for lindex < rindex {
			if tag > nums[rindex] {
				nums[lindex] = nums[rindex]
				lindex++
				break
			}
			rindex--
		}
		for lindex < rindex {
			if tag < nums[lindex] {
				nums[rindex] = nums[lindex]
				rindex--
				break
			}
			lindex++
		}
	}
	nums[lindex] = tag
	// fmt.Printf("nums:%+v\n", nums)
	if kSmallest == lindex {
		return tag
	}
	if lindex > kSmallest {
		return quickselect(left, lindex-1, kSmallest, nums)
	}
	return quickselect(lindex+1, right, kSmallest, nums)
}

func findKthLargest(nums []int, k int) int {
	kSmallest := len(nums) - k
	return quickselect(0, len(nums)-1, kSmallest, nums)
}

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	kMax := findKthLargest(nums, 3)
	fmt.Printf("nums:%+v,max:%d\n", nums, kMax)
}
