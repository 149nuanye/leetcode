package main

import "fmt"

func numDecodingsBack(pre, s string, strHash map[string]string, strList *[]string) {
	if len(s) == 0 {
		if len(pre) == 0 {
			return
		}
		*strList = append(*strList, pre)
		return
	}
	if len(s) == 1 {
		if _, ok := strHash[s]; !ok {
			return
		}
		pre += strHash[s]
		numDecodingsBack(pre, "", strHash, strList)
		return
	}
	tmp := pre
	if _, ok := strHash[string(s[0])]; ok {
		numDecodingsBack(tmp+strHash[string(s[0])], s[1:], strHash, strList)
	}

	if str, ok := strHash[s[:2]]; ok {
		if len(s) == 2 {
			numDecodingsBack(tmp+str, "", strHash, strList)
			return
		}
		numDecodingsBack(tmp+str, s[2:], strHash, strList)
	}
	return
}

func numDecodingsSlow(s string) int {
	strHash := make(map[string]string)
	strHash["1"] = "A"
	strHash["2"] = "B"
	strHash["3"] = "C"
	strHash["4"] = "D"
	strHash["5"] = "E"
	strHash["6"] = "F"
	strHash["7"] = "G"
	strHash["8"] = "H"
	strHash["9"] = "I"
	strHash["10"] = "J"
	strHash["11"] = "K"
	strHash["12"] = "L"
	strHash["13"] = "M"
	strHash["14"] = "N"
	strHash["15"] = "O"
	strHash["16"] = "P"
	strHash["17"] = "Q"
	strHash["18"] = "R"
	strHash["19"] = "S"
	strHash["20"] = "T"
	strHash["21"] = "U"
	strHash["22"] = "V"
	strHash["23"] = "W"
	strHash["24"] = "X"
	strHash["25"] = "Y"
	strHash["26"] = "Z"
	strList := make([]string, 0, len(s))
	numDecodingsBack("", s, strHash, &strList)
	return len(strList)
}

// https://leetcode-cn.com/problems/decode-ways/solution/c-wo-ren-wei-hen-jian-dan-zhi-guan-de-jie-fa-by-pr/
func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	if s[0] == '0' {
		return 0
	}
	pre := 1
	curr := 1
	for i := 1; i < len(s); i++ {
		tmp := curr
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				curr = pre
			} else {
				return 0
			}
		} else if s[i-1] == '1' || (s[i-1] == '2' && s[i] >= '1' && s[i] <= '6') {
			curr = curr + pre
		}
		pre = tmp
	}
	return curr
}

func main() {
	s := "27"
	num := numDecodings(s)
	fmt.Printf("s:%s,num:%d\n", s, num)
}
