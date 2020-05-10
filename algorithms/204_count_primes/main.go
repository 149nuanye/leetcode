package main

import "fmt"

func isPrimes(n int) bool {
	if n == 2 {
		return true
	}
	for index := 2; index*index <= n; index++ {
		if n%index == 0 {
			return false
		}
	}
	return true
}

func countPrimes(n int) int {
	num := 0
	filter := make(map[int]struct{}, n)
	for index := 2; index < n; index++ {
		if _, ok := filter[index]; ok {
			continue
		}
		for j := 2; j*index < n; j++ {
			filter[j*index] = struct{}{}
		}
	}
	for index := 2; index < n; index++ {
		if _, ok := filter[index]; !ok {
			num++
		}
	}
	return num
}

func main() {
	n := 10
	num := countPrimes(n)
	fmt.Printf("n:%d,num:%d\n", n, num)
}
