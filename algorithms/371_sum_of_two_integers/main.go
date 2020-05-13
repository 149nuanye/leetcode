package main

import "fmt"

// https://leetcode-cn.com/problems/sum-of-two-integers/solution/li-yong-wei-cao-zuo-shi-xian-liang-shu-qiu-he-by-p/
func getSum(a int, b int) int {
	for b != 0 {
		tmp := a ^ b
		b = (a & b) << 1
		a = tmp
	}
	return a
}

func main() {
	m := -8
	n := 7
	sum := getSum(m, n)
	fmt.Printf("m:%b,n:%b,sum:%b\n", m, n, sum)
}
