package main

import (
	"fmt"
)

func generate(numRows int) [][]int {
	values := make([][]int, 0, numRows)
	if numRows == 0 {
		return values
	}
	tmp := []int{1}
	if numRows == 1 {
		values = append(values, tmp)
		return values
	}
	values = append(values, tmp)
	for numRows > 1 {
		value := make([]int, 0, len(tmp)+1)
		value = append(value, 1)
		for i := 0; i <= len(tmp)-1; i++ {
			if i+1 <= len(tmp)-1 {
				value = append(value, tmp[i]+tmp[i+1])
			}
		}
		value = append(value, 1)
		tmp = value
		values = append(values, value)
		numRows--
	}
	return values
}

func main() {
	nums := generate(6)
	for index := range nums {
		fmt.Printf("num:%+v\n", nums[index])
	}
}
