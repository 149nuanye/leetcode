package main

import (
	"fmt"
)

func ifPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}
	lindex := 0
	rindex := len(s) - 1
	for lindex < rindex {
		if s[lindex] != s[rindex] {
			return false
		}
		lindex++
		rindex--
	}
	return true
}

func longestPalindrome(s string) string {
	longerStr := ""
	for index := 0; index < len(s); index++ {
		for j := index + 1; j <= len(s); j++ {
			tmpStr := string(s[index:j])
			if len(tmpStr) < len(longerStr) {
				continue
			}
			if ifPalindrome(tmpStr) {
				if len(tmpStr) > len(longerStr) {
					longerStr = tmpStr
				}
			}
		}
	}
	return longerStr
}

func main() {
	str1 := "abadd"
	palindrome := longestPalindrome(str1)
	fmt.Printf("str1:%s, palindrome:%s\n", str1, palindrome)
}
