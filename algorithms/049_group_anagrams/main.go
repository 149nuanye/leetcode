package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	group := make([][]string, 0, len(strs))
	strHash := make(map[string]int)
	for index := range strs {
		tmpList := strings.Split(strs[index], "")
		sort.Strings(tmpList)
		tmpStr := strings.Join(tmpList, "")
		// fmt.Printf("src:%s,tmpStr:%s,group:%+v\n", strs[index], tmpStr, group)
		if num, ok := strHash[tmpStr]; ok {
			group[num] = append(group[num], strs[index])
			continue
		}
		length := len(group)
		strHash[tmpStr] = length
		group = append(group, []string{strs[index]})
	}
	return group
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	group := groupAnagrams(strs)
	for index := range group {
		fmt.Printf("%+v\n", group[index])
	}
}
