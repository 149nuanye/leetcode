package main

import "fmt"

func climbStairs(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	numm := 1
	numn := 2
	for index := 3; index <= n; index++ {
		tmp := numn
		numn = numm + numn
		numm = tmp
	}
	return numn
}

func main() {
	stairs := 3
	num := climbStairs(stairs)
	fmt.Printf("stairs:%d,num:%d\n", stairs, num)
}
