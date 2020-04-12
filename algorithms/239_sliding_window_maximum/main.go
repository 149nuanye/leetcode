package main

import "fmt"

func maxValue(nums []int) int {
	maxV := nums[0]
	for index := 1; index < len(nums); index++ {
		if maxV < nums[index] {
			maxV = nums[index]
		}
	}
	return maxV
}

// 蛮力
func maxSlidingWindowSlow(nums []int, k int) []int {
	maxList := make([]int, 0, len(nums))
	for index := 0; index < len(nums); index++ {
		if len(nums)-index >= k {
			max := maxValue(nums[index : index+k])
			maxList = append(maxList, max)
		}
	}
	return maxList
}

type Queue struct {
	arr []int
}

func (q *Queue) Insert(data int) {
	q.arr = append(q.arr, data)
}

func (q *Queue) DelFront() {
	if len(q.arr) > 1 {
		q.arr = q.arr[1:]
	} else {
		q.arr = make([]int, 0)
	}
}

func (q *Queue) DelTail() {
	if len(q.arr) > 1 {
		q.arr = q.arr[:len(q.arr)-1]
	} else {
		q.arr = make([]int, 0)
	}
}

func (q *Queue) Front() int {
	if len(q.arr) != 0 {
		return q.arr[0]
	}
	return -1
}

func (q *Queue) Tail() int {
	if len(q.arr) != 0 {
		return q.arr[len(q.arr)-1]
	}
	return -1
}

func (q *Queue) Empty() bool {
	if len(q.arr) != 0 {
		return false
	}
	return true
}

// https://leetcode-cn.com/problems/sliding-window-maximum/solution/hua-dong-chuang-kou-zui-da-zhi-by-leetcode-3/
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) <= 1 || k <= 1 {
		return nums
	}
	q := &Queue{}
	var res []int
	for i := 0; i < len(nums); i++ {
		// 如果超过窗口，则删除
		if !q.Empty() && i >= q.Front()+k {
			q.DelFront()
		}
		for !q.Empty() && nums[i] >= nums[q.Tail()] {
			q.DelTail()
		}
		q.Insert(i)

		if i+1 >= k {
			res = append(res, nums[q.Front()])
		}
	}
	return res
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	fmt.Printf("nums:%+v\n", nums)
	maxList := maxSlidingWindow(nums, 3)
	fmt.Printf("maxList:%+v\n", maxList)
}
