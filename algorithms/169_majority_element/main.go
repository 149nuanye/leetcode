package main

import "fmt"

func majorityElementHash(nums []int) int {
	valueHash := make(map[int]int)
	for index := range nums {
		if num, ok := valueHash[nums[index]]; ok {
			valueHash[nums[index]] = num + 1
		} else {
			valueHash[nums[index]] = 1
		}
	}
	majority := 0
	for index := range nums {
		if num, ok := valueHash[nums[index]]; ok {
			if num > len(nums)/2 {
				return nums[index]
			}
		}
	}
	return majority
}

// 利用 Boyer-Moore 投票算法(这个需要保证存在众数)
/* 寻找数组中超过一半的数字，这意味着数组中其他数字出现次数的总和都是比不上这个数字出现的次数。
即如果把 该众数记为 +1 ，把其他数记为 −1 ，将它们全部加起来，和是大于 0 的。*/
func majorityElement(nums []int) int {
	count := 1
	majority := nums[0]
	for index := 1; index < len(nums); index++ {
		if count == 0 {
			majority = nums[index]
			count = 1
			continue
		}
		if majority == nums[index] {
			count++
		} else {
			count--
		}
	}
	return majority
}

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Printf("nums:%+v\n", nums)
	majority := majorityElement(nums)
	fmt.Printf("majority:%+v\n", majority)
}
