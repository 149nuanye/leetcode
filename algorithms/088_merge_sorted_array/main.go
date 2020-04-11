package main

import "fmt"

// 从数组后面插入
func merge(nums1 []int, m int, nums2 []int, n int) {
	index1 := m - 1
	index2 := n - 1
	newIndex := m + n - 1
	for index1 >= 0 && index2 >= 0 {
		if nums1[index1] > nums2[index2] {
			nums1[newIndex] = nums1[index1]
			index1--
		} else {
			nums1[newIndex] = nums2[index2]
			index2--
		}
		newIndex--
	}
	for index1 >= 0 {
		nums1[newIndex] = nums1[index1]
		index1--
		newIndex--
	}
	for index2 >= 0 {
		nums1[newIndex] = nums2[index2]
		index2--
		newIndex--
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
	fmt.Printf("nums1:%+v\n", nums1)
}
