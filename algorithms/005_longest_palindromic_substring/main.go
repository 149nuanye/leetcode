package main

import (
	"fmt"
)

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	longestStr := make([]byte, 0, len(s))
	lStr := make([]byte, 0, (len(s)+1)/2)
	rStr := make([]byte, 0, (len(s)+1)/2)
	begin := true
	for lindex := 0; lindex < len(s); lindex++ {
		if begin {
			lStr = lStr[:0]
			rStr = rStr[:0]
			rStr = append(rStr, s[lindex])
			begin = false
		}
		if lindex+1 < len(s) {
			if s[lindex] == s[lindex+1] {
				rStr = append(rStr, s[lindex+1])
				continue
			}
		}
		begin = true

		i := lindex - len(rStr)
		j := lindex + 1
		for i >= 0 && j < len(s) {
			if s[i] == s[j] {
				rStr = append(rStr, s[j])
				lStr = append(lStr, s[i])
			} else {
				break
			}
			i--
			j++
		}
		// fmt.Printf("lStr:%s,rStr:%s,longestStr:%s\n", rStr, longestStr)
		if len(lStr)+len(rStr) > len(longestStr) {
			longestStr = longestStr[:0]
			for index := len(lStr) - 1; index >= 0; index-- {
				longestStr = append(longestStr, lStr[index])
			}
			longestStr = append(longestStr, rStr...)
		}
	}
	return string(longestStr)
}

func main() {
	str1 := "abadd"
	palindrome := longestPalindrome(str1)
	fmt.Printf("str1:%s,palindrome:%s\n", str1, palindrome)
}
