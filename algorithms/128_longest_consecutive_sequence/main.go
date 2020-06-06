package main

import "fmt"

func longestConsecutive(nums []int) int {
	longest := 0
	if len(nums) == 0 {
		return longest
	}
	numHash := make(map[int]struct{})
	for index := range nums {
		numHash[nums[index]] = struct{}{}
	}
	for _, num := range nums {
		tmpLongest := 1
		if _, ok := numHash[num-1]; ok {
			continue
		}
		for {
			num++
			if _, ok := numHash[num]; !ok {
				break
			}
			tmpLongest++
		}
		if tmpLongest > longest {
			longest = tmpLongest
		}
	}
	return longest
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	longest := longestConsecutive(nums)
	fmt.Printf("nums:%+v,longest:%d\n", nums, longest)
}
