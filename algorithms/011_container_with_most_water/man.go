package main

import "fmt"

func maxArea(height []int) int {
	area := 0
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			tmp := height[i]
			if height[i] > height[j] {
				tmp = height[j]
			}
			if area < tmp*(j-i) {
				area = tmp * (j - i)
			}
		}
	}
	return area
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	area := maxArea(height)
	fmt.Printf("height:%+v,area:%+v\n", height, area)
}
