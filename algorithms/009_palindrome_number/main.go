package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	newNum := 0
	srcNum := x
	for x > 0 {
		newNum = newNum*10 + x%10
		x = x / 10
	}
	if newNum != srcNum {
		return false
	}
	return true
}

func main() {
	num := 1215
	palindrome := isPalindrome(num)
	fmt.Printf("num:%d, palindrome:%+v\n", num, palindrome)
}
