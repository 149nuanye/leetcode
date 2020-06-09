package main

import "fmt"

func myPowSlow(x float64, n int) float64 {
	tag := false
	if n < 0{
		tag = true
		n = -n
	}
	var value float64 =1.0
	for i:=0; i<n; i++{
		value *=x
	}
	if tag {
		value=1/value
	}
	return value
}

// https://leetcode-cn.com/problems/powx-n/solution/powx-n-by-leetcode-solution/
func quickMul(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	y := quickMul(x,n/2)
	if n % 2 ==0 {
		return y*y
	} 
	return y*y*x
}

func myPow(x float64, n int) float64 {
	if n < 0{
		return 1/quickMul(x,n)
	}
	return quickMul(x,n)
}

func main(){
	x:=2.00000
	n:= -2
	value:=myPow(x,n)
	fmt.Printf("x:%+v,n:%+v,value:%+v\n",x,n,value)
}