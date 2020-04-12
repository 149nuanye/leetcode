package main

import "fmt"

func reverse(x int) int {
	newValue := 0
	maxValue := 1 << 31
	if x < -(maxValue) || x > (maxValue-1) {
		return 0
	}
	tag := false
	if x < 0 {
		tag = true
		x = -x
	}
	for x > 0 {
		tmp := x % 10
		x = x / 10
		newValue = newValue*10 + tmp
		if newValue > maxValue-1 {
			return 0
		}
	}
	if tag {
		return -newValue
	}
	return newValue
}

func main() {
	num := -810
	fmt.Printf("num:%d\n", num)
	reverseValue := reverse(num)
	fmt.Printf("reverseValue:%d\n", reverseValue)
}
