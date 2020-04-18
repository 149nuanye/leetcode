package main

import (
	"fmt"
	"strconv"
)

// Stack stack
type Stack struct {
	List []string
}

// New new stack
func New() *Stack {
	st := &Stack{List: make([]string, 0, 10)}
	return st
}

// Push push
func (st *Stack) Push(value string) {
	st.List = append(st.List, value)
}

// Pop pop
func (st *Stack) Pop() string {
	value := ""
	if len(st.List) == 0 {
		return value
	}
	value = st.List[len(st.List)-1]
	st.List = st.List[:len(st.List)-1]
	return value
}

// Len len
func (st *Stack) Len() int {
	return len(st.List)
}

func evalRPN(tokens []string) int {
	if len(tokens) == 1 {
		b, _ := strconv.Atoi(tokens[0])
		return b
	}

	value := 0
	stack := New()
	for index := range tokens {
		a := 0
		b := 0
		switch tokens[index] {
		case "+", "-", "*", "/":
			a, _ = strconv.Atoi(stack.Pop())
			b, _ = strconv.Atoi(stack.Pop())
			break
		default:
			stack.Push(tokens[index])
			continue
		}
		if tokens[index] == "+" {
			value = a + b
		} else if tokens[index] == "-" {
			value = b - a
		} else if tokens[index] == "*" {
			value = a * b
		} else if tokens[index] == "/" {
			value = b / a
		}
		if len(tokens)-1 != index {
			stack.Push(strconv.Itoa(value))
		}
	}
	return value
}

func main() {
	tokens := []string{"4", "13", "5", "/", "+"}
	value := evalRPN(tokens)
	fmt.Printf("value:%d\n", value)
}
