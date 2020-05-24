package main

import "fmt"

func isPalind(s string) bool {
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

func palindrome(strList []string, s string, palinList *[][]string) {
	if len(s) <= 1 {
		if len(s) == 1 {
			strList = append(strList, s)
		}
		*palinList = append(*palinList, strList)
		return
	}
	for index := 0; index < len(s); index++ {
		tmpList := make([]string, 0, len(strList)+4)
		tmpList = append(tmpList, strList...)
		if isPalind(string(s[:index+1])) {
			tmpList = append(tmpList, string(s[:index+1]))
			if index <= len(s)-1 {
				palindrome(tmpList, string(s[index+1:]), palinList)
			}
		}
	}
}

func partition(s string) [][]string {
	palinList := make([][]string, 0, 10)
	palindrome(nil, s, &palinList)
	return palinList
}

func main() {
	s := "aa"
	strList := partition(s)
	for index := range strList {
		fmt.Printf("value:%+v\n", strList[index])
	}
}
