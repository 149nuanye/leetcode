package main

import "fmt"

//  时间复杂度：O(m + n)，找顺序
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	medianValue := 0.0
	double := false
	medianIndex := 0
	if (len(nums1)+len(nums2))%2 == 0 {
		medianIndex = (len(nums1) + len(nums2)) / 2
		double = true
	} else {
		medianIndex = (len(nums1) + len(nums2) + 1) / 2
	}
	num := 0
	index1, index2 := 0, 0
	for index1 < len(nums1) || index2 < len(nums2) {
		tmp := 0
		if index1 >= len(nums1) {
			tmp = nums2[index2]
			index2++
		} else if index2 >= len(nums2) {
			tmp = nums1[index1]
			index1++
		} else {
			if nums1[index1] <= nums2[index2] {
				tmp = nums1[index1]
				index1++
			} else {
				tmp = nums2[index2]
				index2++
			}
		}
		num++
		if num > medianIndex+1 {
			break
		}
		if num == medianIndex {
			medianValue = float64(tmp)
		}
		if double && num == medianIndex+1 {
			medianValue = (medianValue + float64(tmp)) / 2
		}
	}
	return medianValue
}

func main() {
	nums1 := []int{5, 6}
	nums2 := []int{3, 4}
	fmt.Printf("nums1:%+v\n", nums1)
	fmt.Printf("nums2:%+v\n", nums2)
	medianValue := findMedianSortedArrays(nums1, nums2)
	fmt.Printf("medianValue:%+v\n", medianValue)
}
