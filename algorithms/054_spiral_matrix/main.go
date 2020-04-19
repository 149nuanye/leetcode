package main

import "fmt"

// https://leetcode-cn.com/problems/spiral-matrix/solution/luo-xuan-ju-zhen-by-leetcode/
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	n := len(matrix)
	m := len(matrix[0])
	order := make([]int, 0, n*m)
	nindex := 0
	mindex := 0
	for nindex < n && mindex < m {
		for i := mindex; i < m; i++ {
			order = append(order, matrix[nindex][i])
		}
		for j := nindex + 1; j < n; j++ {
			order = append(order, matrix[j][m-1])
		}
		if n-nindex > 1 {
			for i := m - 2; i >= mindex; i-- {
				order = append(order, matrix[n-1][i])
			}
		}
		if m-mindex > 1 {
			for j := n - 2; j > nindex; j-- {
				order = append(order, matrix[j][mindex])
			}
		}
		nindex++
		mindex++
		m--
		n--
	}
	return order
}

func main() {
	matrix := [][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}
	// matrix := [][]int{[]int{1,2,3}}
	order := spiralOrder(matrix)
	fmt.Printf("order:%+v\n", order)
}
