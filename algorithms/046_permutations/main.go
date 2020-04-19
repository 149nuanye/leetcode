package main

import "fmt"
// https://studygolang.com/articles/19971
func permuteImp(preNums,nums []int,per *[][]int){
	if preNums == nil {
		preNums = make([]int,0,len(nums))
	}
	if len(nums)==1{
		tmp := make([]int,len(preNums)+1)
		copy(tmp,preNums)
		tmp[len(preNums)] = nums[0]
		*per = append(*per,tmp)
		return 
	}
	for index := 0;index<len(nums);index++ {
		tmp:=make([]int,0,len(nums)-2)
		if index > 0{
			tmp = append(tmp,nums[:index]...)
		} 
		if index < len(nums)-1{
			tmp = append(tmp,nums[index+1:]...)
		}
		preNums = append(preNums,nums[index])
		// fmt.Printf("index:%d,nums:%+v,tmp:%+v,preNums:%+v\n", index, nums,tmp,preNums)
		permuteImp(preNums,tmp,per)
		preNums = preNums[:len(preNums)-1]
		// fmt.Printf("end index:%d,nums:%+v,tmp:%+v,preNums:%+v\n", index, nums,tmp,preNums)
	}
}

func permute(nums []int) [][]int {
	length := len(nums) 
	per := make([][]int,0,length*(length-1)+1)
	permuteImp(nil,nums,&per)
	return per
}

func main() {
	nums := []int{1, 2}
	values := permute(nums)
	for index := range values {
		fmt.Printf("index:%d,value:%+v\n", index, values[index])
	}
}
