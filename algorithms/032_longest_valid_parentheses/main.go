package main

import "fmt"

// Stack stack
type Stack struct {
	List []int
}

// New new stack
func New() *Stack {
	st := &Stack{List: make([]int, 0, 10)}
	return st
}

// Push push
func (st *Stack) Push(value int) {
	st.List = append(st.List, value)
}

// Pop pop
func (st *Stack) Pop() int {
	var value int
	if len(st.List) == 0 {
		return value
	}
	value = st.List[len(st.List)-1]
	st.List = st.List[:len(st.List)-1]
	return value
}

// First first
func (st *Stack) First() int {
	if len(st.List) > 0 {
		return st.List[len(st.List)-1]
	}
	return -1
}

// Len len
func (st *Stack) Len() int {
	return len(st.List)
}

func longestValidParentheses(s string) int {
	max := 0
	stack := New()
	for index := range s {
		if s[index] == '(' || stack.Len() == 0 {
			stack.Push(index)
			continue
		}
		first := stack.First()
		if s[first] == ')' {
			stack.Push(index)
			continue
		}
		stack.Pop()
		num := index - stack.First()
		if num > max {
			max = num
		}
	}
	return max
}

func main() {
	str := ")()())"
	num := longestValidParentheses(str)
	fmt.Printf("str:%s, num:%d\n", str, num)
}
