package main

import (
	"fmt"
)

func reverse(x int) int {
	tag, tmp := 1, x
	if x < 0 {
		tag = -1
		tmp = -x
	}
	intArr := make([]int, 0, 10)
	for tmp != 0 {
		intArr = append(intArr, tmp%10)
		tmp = tmp / 10
	}
	tag2 := 0
	for index := range intArr {
		if tag2 == 0 && intArr[index] == 0 {
			continue
		}
		tag2 = 1
		tmp = 10*tmp + intArr[index]
	}
	max := 1 << 31
	if tag < 0 {
		if -tmp < -max {
			return 0
		}
		return -tmp
	}
	if tmp > max {
		return 0
	}
	return tmp
}
func main() {
	val := reverse(-7483648)
	fmt.Printf("val:%d\n", val)
}
