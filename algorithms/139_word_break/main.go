package main

import "fmt"

func wordBreakMem(start int, s string, dictMap map[string]struct{}, memory map[int]bool) bool {
	// fmt.Printf("start:%d, s:%s,dictMap:%+v\n", start, s, dictMap)
	if _, ok := memory[start]; ok {
		return memory[start]
	}
	for index := start; index <= len(s); index++ {
		if _, ok := dictMap[s[start:index]]; ok {
			if index == len(s) || wordBreakMem(index, s, dictMap, memory) {
				memory[start] = true
				return true
			}
		}
	}
	memory[start] = false
	return false
}
func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}
	dictMap := make(map[string]struct{})
	for index := range wordDict {
		dictMap[wordDict[index]] = struct{}{}
	}
	memory := make(map[int]bool)
	return wordBreakMem(0, s, dictMap, memory)
}

func main() {
	s := "aaaaaa"
	wordDict := []string{"aaaa", "aaa"}
	wb := wordBreak(s, wordDict)
	fmt.Printf("s:%s,wordDict:%+v,wb:%+v\n", s, wordDict, wb)
}
