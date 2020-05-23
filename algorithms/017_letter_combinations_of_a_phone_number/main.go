package main

import "fmt"

func combnations(pre []string, digits string, letterMap map[byte][]string) []string {
	strList := make([]string, 0, 100)
	if len(digits) == 0 {
		if len(pre) == 0 {
			return strList
		}
		return []string{pre}
	}
	digit := digits[0]
	mapStrs := letterMap[digit]
	digits = string(digits[1:])
	for index := range mapStrs {
		temp := pre
		pre = pre + mapStrs[index]
		strList = append(strList, combnations(pre, digits, letterMap)...)
		pre = temp
	}
	return strList
}

func letterCombinations(digits string) []string {
	letterMap := make(map[byte][]string)
	letterMap['2'] = []string{"a", "b", "c"}
	letterMap['3'] = []string{"d", "e", "f"}
	letterMap['4'] = []string{"g", "h", "i"}
	letterMap['5'] = []string{"j", "k", "l"}
	letterMap['6'] = []string{"m", "n", "o"}
	letterMap['7'] = []string{"p", "q", "r", "s"}
	letterMap['8'] = []string{"t", "u", "v"}
	letterMap['9'] = []string{"w", "x", "y", "z"}
	return combnations("", digits, letterMap)
}

func main() {
	str := "23"
	strList := letterCombinations(str)
	fmt.Printf("str:%s\n", str)
	for index := range strList {
		fmt.Printf("index:%d,value:%s\n", index, strList[index])
	}
}
