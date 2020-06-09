package main

import "fmt"

func numSquares(n int) int {
	squareNums := make([]int, 0, n)
	for i := 1; i*i <= n; i++ {
		squareNums = append(squareNums, i*i)
	}
	fmt.Printf("squareNums:%+v\n", squareNums)
	dp := make([]int, 0, n)
	dp = append(dp, 0)
	for i := 1; i <= n; i++ {
		tmp := i
		for _, j := range squareNums {
			if i < j {
				break
			}
			if dp[i-j]+1 < tmp {
				tmp = dp[i-j] + 1
			}
		}
		dp = append(dp, tmp)
		fmt.Printf("i:%+v,tmp:%d,dp:%+v\n", i, tmp, dp)
	}
	return dp[len(dp)-1]
}

func main() {
	n := 4
	num := numSquares(n)
	fmt.Printf("n:%d,num:%d\n", n, num)
}
