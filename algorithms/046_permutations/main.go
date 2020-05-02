package main

import "fmt"

// 全排列问题，比如打印1-9的共9个字母的全排列。先输出1，
// 然后是2-9的全排列，输出2，然后1-9中去除2的全排列。
// 于是很自然想到递归的方法。写出伪代码：
// permutaion(prefix, set) {
//     if set 为空
//         print prefix

//     for s in set:
//         permuetaion(prefix+s, set - s)
// }
// https://studygolang.com/articles/19971
func permuteImp(preNums, nums []int, per *[][]int) {
	if preNums == nil {
		preNums = make([]int, 0, len(nums))
	}
	if len(nums) == 1 {
		tmp := make([]int, len(preNums)+1)
		copy(tmp, preNums)
		tmp[len(preNums)] = nums[0]
		*per = append(*per, tmp)
		return
	}
	for index := 0; index < len(nums); index++ {
		tmp := make([]int, 0, len(nums)-1)
		if index > 0 {
			tmp = append(tmp, nums[:index]...)
		}
		if index < len(nums)-1 {
			tmp = append(tmp, nums[index+1:]...)
		}
		preNums = append(preNums, nums[index])
		permuteImp(preNums, tmp, per)
		preNums = preNums[:len(preNums)-1]
	}
}

func permute(nums []int) [][]int {
	length := len(nums)
	per := make([][]int, 0, length*(length-1)+1)
	permuteImp(nil, nums, &per)
	return per
}

func main() {
	nums := []int{1, 2}
	values := permute(nums)
	for index := range values {
		fmt.Printf("index:%d,value:%+v\n", index, values[index])
	}
}
