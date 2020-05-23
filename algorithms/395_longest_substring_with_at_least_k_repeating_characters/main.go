package main

import "fmt"

// https://leetcode-cn.com/problems/longest-substring-with-at-least-k-repeating-characters/solution/duo-lu-fen-zhi-de-di-gui-fang-fa-zhi-xing-he-nei-c/
func longestSubstring(s string, k int) int {
	maxLength := 0
	if len(s) < k {
		return maxLength
	}
	numStatic := make(map[byte]int)
	for index := range s {
		numStatic[s[index]] += 1
	}
	// fmt.Printf("s:%s,k:%d,numStatic:%+v\n", s, k, numStatic)
	splite := false
	for index := range s {
		if numStatic[s[index]] < k {
			max := longestSubstring(string(s[:index]), k)
			if maxLength < max {
				maxLength = max
			}
			if index+1 <= len(s)-1 {
				max = longestSubstring(string(s[index+1:]), k)
				if maxLength < max {
					maxLength = max
				}
			}
			splite = true
			break
		}
	}
	if !splite {
		return len(s)
	}
	// fmt.Printf("s:%s,k:%d,numStatic:%+v,maxLength:%d\n", s, k, numStatic, maxLength)
	return maxLength
}

func main() {
	s := "aacabb"
	k := 2
	maxLength := longestSubstring(s, k)
	fmt.Printf("s:%s,k:%d,maxLength:%d\n", s, k, maxLength)
}
