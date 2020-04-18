package main

import "fmt"

func mySqrt(x int) int {
	if x < 2 {
		return x
	}
	lindex := 2
	rindex := x
	num := 0
	for lindex <= rindex {
		// fmt.Printf("lindex:%d, rindex:%d\n", lindex, rindex)
		num = (lindex + rindex) / 2
		value := num * num
		if value > x {
			rindex = num - 1
		} else if value < x {
			lindex = num + 1
		} else {
			return num
		}
	}
	if num*num > x {
		return num - 1
	}
	return num
}

func main() {
	num := 2
	value := mySqrt(num)
	fmt.Printf("num:%d\nvalue:%d\n", num, value)
}
