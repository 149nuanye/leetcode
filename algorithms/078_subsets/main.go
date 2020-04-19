package main

import "fmt"

// https://leetcode-cn.com/problems/subsets/solution/zi-ji-by-leetcode/
func subsets(nums []int) [][]int {
	sets := make([][]int, 0, 10)
	for index := range nums {
		length := len(sets)
		tmpSets := make([][]int, 0, 10)
		for i := 0; i < length; i++ {
			tmp := make([]int, len(sets[i]))
			copy(tmp, sets[i])
			tmpSets = append(tmpSets, append(tmp, nums[index]))
		}
		sets = append(sets, []int{nums[index]})
		if len(tmpSets) > 0 {
			sets = append(sets, tmpSets...)
		}
	}
	sets = append(sets, make([]int, 0))
	return sets
}

func main() {
	nums := []int{9, 0, 3, 5, 7}
	sets := subsets(nums)
	for index := range sets {
		fmt.Printf("index:%d,data:%+v\n", index, sets[index])
	}
}
