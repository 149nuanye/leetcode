package main

import (
	"fmt"
	"strconv"
)

func compare(a, b int) bool {
	str1 := strconv.Itoa(a)
	str2 := strconv.Itoa(b)
	value1 := str1 + str2
	value2 := str2 + str1
	num1, _ := strconv.Atoi(value1)
	num2, _ := strconv.Atoi(value2)
	if num1 > num2 {
		return true
	}
	return false
}

func sortNums(nums []int) {
	if len(nums) < 2 {
		return
	}
	tag := nums[0]
	lindex := 0
	rindex := len(nums) - 1
	for lindex < rindex {
		for rindex > lindex {
			if compare(nums[rindex], tag) {
				rindex--
				continue
			}
			nums[lindex] = nums[rindex]
			lindex++
			break
		}
		for rindex > lindex {
			if !compare(nums[lindex], tag) {
				lindex++
				continue
			}
			nums[rindex] = nums[lindex]
			rindex--
			break
		}
	}
	nums[lindex] = tag
	sortNums(nums[:lindex])
	if lindex < len(nums)-1 {
		sortNums(nums[lindex+1:])
	}
}

func largestNumber(nums []int) string {
	sortNums(nums)
	// fmt.Printf("sort nums:%+v\n", nums)
	strNum := ""
	for i := len(nums) - 1; i >= 0; i-- {
		str := strconv.Itoa(nums[i])
		if str == "0" && strNum == "0" {
			continue
		}
		strNum = strNum + str
	}
	return strNum
}

func main() {
	nums := []int{3, 30, 34, 5, 9}
	fmt.Printf("nums:%+v\n", nums)
	largeNum := largestNumber(nums)
	fmt.Printf("largeNum:%+v\n", largeNum)
}
