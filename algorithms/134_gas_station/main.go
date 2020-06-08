package main

import "fmt"

// https://leetcode-cn.com/problems/gas-station/solution/jia-you-zhan-by-leetcode/
func canCompleteCircuit(gas []int, cost []int) int {
	totalGas := 0
	curGas := 0
	startIndex := 0
	for index := range gas {
		totalGas += gas[index] - cost[index]
		curGas += gas[index] - cost[index]
		if curGas < 0 {
			startIndex = index + 1
			curGas = 0
		}
	}
	if totalGas < 0 {
		return -1
	}
	return startIndex
}

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	value := canCompleteCircuit(gas, cost)
	fmt.Printf("gas:%+v,cost:%+v,value:%d\n", gas, cost, value)
}
