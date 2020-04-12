package main

import "fmt"

func myAtoi(str string) int {
	newValue := 0
	tag := false
	frist := true
	minValue := -1 << 31
	maxValue := (1<<31 - 1)
	for index := 0; index < len(str); index++ {
		if newValue > maxValue {
			break
		}
		if frist {
			frist = false
			if str[index] == ' ' {
				frist = true
				continue
			} else if str[index] == '+' {
				continue
			} else if str[index] == '-' {
				tag = true
				continue
			}
		}
		if str[index] >= '0' && str[index] <= '9' {
			newValue = newValue*10 + int(str[index]-'0')
		} else {
			break
		}
	}
	fmt.Printf("tag:%+v,newValue:%d\n", tag, newValue)

	if tag {
		newValue = -newValue
	}
	if newValue < minValue {
		return minValue
	} else if newValue > maxValue {
		return maxValue
	}
	return newValue
}

func main() {
	str := "9223372036854775808"
	fmt.Printf("str:%s\n", str)
	newValue := myAtoi(str)
	fmt.Printf("newValue:%d\n", newValue)
}
