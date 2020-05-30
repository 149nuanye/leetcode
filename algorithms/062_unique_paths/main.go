package main

import "fmt"

func uniquePathsMath(m int, n int) int {
	if n > m {
		n, m = m, n
	}
	var num1 int64 = int64(m + n - 2)
	var num2 int64 = int64(n - 1)
	var value1 int64 = 1
	var value2 int64 = 1
	for num2 > 0 {
		// fmt.Printf("num1:%d,num2:%d,value1:%d,value2:%d\n", num1, num2, value1, value2)
		value1 *= num1
		value2 *= num2
		if value1%value2 == 0 {
			value1 = value1 / value2
			value2 = 1
		}
		num1--
		num2--
	}
	return int(value1 / value2)
}

func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			tmp := 0
			if i == 0 || j == 0 {
				tmp = 1
			} else {
				tmp = dp[i-1][j] + dp[i][j-1]
			}
			dp[i] = append(dp[i], tmp)
		}
	}
	return dp[m-1][n-1]
}

func main() {
	m := 13
	n := 23
	path := uniquePaths(m, n)
	fmt.Printf("m:%d,n:%d,path:%d\n", m, n, path)
}
