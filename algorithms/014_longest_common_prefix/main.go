package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	num := len(strs[0])
	shortIndex := 0
	for index := 1; index < len(strs); index++ {
		if len(strs[index]) < num {
			num = len(strs[index])
			shortIndex = index
		}
	}
	commonPrefix := strs[shortIndex]

	for index := 0; index < len(strs); index++ {
		i := 0
		j := 0
		if index == shortIndex {
			continue
		}
		for i < len(commonPrefix) && j < len(strs[index]) {
			if commonPrefix[i] != strs[index][j] {
				break
			}
			i++
			j++
		}
		if i == 0 {
			return ""
		}
		if i < len(commonPrefix) {
			commonPrefix = commonPrefix[:i]
		}
	}
	return commonPrefix
}

func main() {
	strArr := []string{"aa", "ab"}
	fmt.Printf("strArr:%+v\n", strArr)
	commonPrefix := longestCommonPrefix(strArr)
	fmt.Printf("commonPrefix:%s\n", commonPrefix)
}
