package main

import (
	"fmt"
)

func demoQuick() {
	nums := []int{10, 2, 5, 7, 20}
	fmt.Printf("nums:%+v\n", nums)
	quickSort(nums)
	fmt.Printf("after sort nums:%+v\n", nums)
}
func main() {
	demoQuick()
}
