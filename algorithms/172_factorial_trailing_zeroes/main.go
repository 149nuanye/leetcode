package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/factorial-trailing-zeroes/solution/xiang-xi-tong-su-de-si-lu-fen-xi-by-windliang-3/
func trailingZeroesSlow(n int) int {
	num := 0
	for n > 0 {
		tmp := n
		for tmp >= 5 {
			if tmp%5 == 0 {
				num++
				tmp = tmp / 5
				continue
			}
			break
		}
		n--
	}
	return num
}

func trailingZeroesFast(n int) int {
	num := 0
	for index := 5; index <= n; index += 5 {
		tmp := index
		for tmp >= 5 {
			if tmp%5 == 0 {
				num++
				tmp = tmp / 5
				continue
			}
			break
		}
	}
	return num
}

func trailingZeroes(n int) int {
	num := 0
	for n > 0 {
		num += n / 5
		n = n / 5
	}
	return num
}

func main() {
	n := 10
	num := trailingZeroes(n)
	fmt.Printf("n:%d,num:%d\n", n, num)
}
