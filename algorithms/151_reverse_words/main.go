package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	values := strings.Split(s, " ")
	newStr := ""
	for index := len(values) - 1; index >= 0; index-- {
		if len(values[index]) == 0 {
			continue
		}
		if len(newStr) == 0 {
			newStr = strings.TrimSpace(values[index])
		} else {
			newStr += " " + values[index]
		}
	}
	return newStr
}

func main() {
	str := "hello world"
	fmt.Printf("str:%s\n", str)
	newStr := reverseWords(str)
	fmt.Printf("newStr:%s\n", newStr)
}
