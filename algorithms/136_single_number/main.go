package main

import "fmt"

func singleNumber(nums []int) int {
	value := 0
	for index := range nums {
		value = value ^ nums[index]
	}
	return value
}

func main() {
	nums := []int{4, 1, 2, 1, 2}
	value := singleNumber(nums)
	fmt.Printf("value:%d\n", value)
}
