package main

import "fmt"

// Heap heap
type Heap struct {
	Value []byte
}

// NewHeap new heap
func NewHeap() *Heap {
	h := &Heap{Value: make([]byte, 0, 100)}
	return h
}

//Push push
func (h *Heap) Push(v byte) int {
	h.Value = append(h.Value, v)
	return len(h.Value)
}

//Pop pop
func (h *Heap) Pop() byte {
	if len(h.Value) == 0 {
		return 0
	}
	tmp := h.Value[len(h.Value)-1]
	h.Value = h.Value[:len(h.Value)-1]
	return tmp
}

//Length length
func (h *Heap) Length() int {
	return len(h.Value)
}

func isValid(s string) bool {
	heap := NewHeap()
	for index := range s {
		if s[index] == '(' || s[index] == '[' || s[index] == '{' {
			heap.Push(s[index])
		} else if s[index] == ')' {
			tmp := heap.Pop()
			if tmp != '(' {
				return false
			}
		} else if s[index] == ']' {
			tmp := heap.Pop()
			if tmp != '[' {
				return false
			}
		} else if s[index] == '}' {
			tmp := heap.Pop()
			if tmp != '{' {
				return false
			}
		} else {
			return false
		}
	}
	if heap.Length() > 0 {
		return false
	}
	return true
}

func main() {
	str := "()[]{}"
	fmt.Printf("str:%s\n", str)
	fmt.Printf("isValid:%+v\n", isValid(str))
}
