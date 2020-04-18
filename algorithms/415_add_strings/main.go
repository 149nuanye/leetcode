package main

import "fmt"

func addStrings(num1 string, num2 string) string {
	value := make([]byte, 0, 5100)
	i := len(num1) - 1
	j := len(num2) - 1
	var tag uint8
	for i >= 0 && j >= 0 {
		tmp := num1[i] + (num2[j] - '0') + tag
		if tmp > '9' {
			tmp = tmp - 10
			tag = 1
		} else {
			tag = 0
		}
		value = append(value, tmp)
		i--
		j--
	}
	for i >= 0 {
		tmp := num1[i] + tag
		if tmp > '9' {
			tmp = tmp - 10
			tag = 1
		} else {
			tag = 0
		}
		value = append(value, tmp)
		i--
	}
	for j >= 0 {
		tmp := num2[j] + tag
		if tmp > '9' {
			tmp = tmp - 10
			tag = 1
		} else {
			tag = 0
		}
		value = append(value, tmp)
		j--
	}
	if tag == 1 {
		value = append(value, '1')
	}
	lindex := 0
	rindex := len(value) - 1
	for lindex < rindex {
		value[lindex], value[rindex] = value[rindex], value[lindex]
		lindex++
		rindex--
	}
	return string(value)
}

func main() {
	num1 := "17"
	num2 := "345"
	sumNum := addStrings(num1, num2)
	fmt.Printf("num1:%s\nnum2:%s\nsum:%s\n", num1, num2, sumNum)
}
