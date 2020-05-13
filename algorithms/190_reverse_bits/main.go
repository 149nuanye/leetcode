package main

import (
	"fmt"
)

func reverseBits(num uint32) uint32 {
	var newValue uint32
	for index := 31; index >= 0; index-- {
		newValue += (num & 1 << index)
		num = num >> 1
	}
	return newValue
}

func main() {
	var m uint32 = 43261596
	n := reverseBits(m)
	fmt.Printf("m:%d,n:%d\n", m, n)
}
