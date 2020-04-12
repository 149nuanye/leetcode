package main

import "fmt"

func reverseString(s []byte) {
	lindex := 0
	rindex := len(s) - 1
	for lindex < rindex {
		tmp := s[lindex]
		s[lindex] = s[rindex]
		s[rindex] = tmp
		// s[rindex], s[lindex] = s[lindex], s[rindex]
		lindex++
		rindex--
	}
}

func main() {
	str1 := []byte{'h', 'e', 'l', 'l', 'o'}
	fmt.Printf("str1:%s\n", str1)
	reverseString(str1)
	fmt.Printf("str1:%s\n", str1)
}
