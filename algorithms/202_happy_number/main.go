package main

import "fmt"

func isHappy(n int) bool {
	filter := make(map[int]struct{})
	for n > 0 {
		fmt.Printf("n:%d\n", n)
		if _, ok := filter[n]; ok {
			return false
		}
		if n == 1 {
			return true
		}
		filter[n] = struct{}{}
		if n < 10 {
			n = n * n
		}
		tmp := 0
		for n > 0 {
			i := n % 10
			n = n / 10
			tmp += i * i
		}
		n = tmp
	}
	return false
}

func main() {
	n := 19
	happy := isHappy(n)
	fmt.Printf("n:%d,happy:%+v\n", n, happy)
}
