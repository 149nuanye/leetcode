package main

type MinStack struct {
	Data    []int
	MinList []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{Data: make([]int, 0, 100), MinList: make([]int, 0, 100)}
}

func (this *MinStack) Push(x int) {
	if len(this.MinList) == 0 || x <= this.GetMin() {
		this.MinList = append(this.MinList, x)
	}
	this.Data = append(this.Data, x)
}

func (this *MinStack) Pop() {
	if len(this.Data) == 0 {
		return
	}
	x := this.Data[len(this.Data)-1]
	if x == this.MinList[len(this.MinList)-1] {
		this.MinList = this.MinList[:len(this.MinList)-1]
	}
	this.Data = this.Data[:len(this.Data)-1]
}

func (this *MinStack) Top() int {
	if len(this.Data) == 0 {
		return -1
	}
	return this.Data[len(this.Data)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.MinList) == 0 {
		return -1
	}
	return this.MinList[len(this.MinList)-1]
}

func main() {

}
