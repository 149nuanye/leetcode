package main

import "fmt"

// Queue queue
type Queue struct {
	Data []byte
}

// NewQueue new queue
func NewQueue() *Queue {
	queue := &Queue{Data: make([]byte, 0, 100)}
	return queue
}

func (q *Queue) Push(b byte) {
	q.Data = append(q.Data, b)
}

func (q *Queue) Pop() byte {
	b := q.Data[0]
	q.Data = q.Data[1:]
	return b
}

func (q *Queue) Length() int {
	return len(q.Data)
}

func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	if len(s) == 0 {
		return maxLength
	}
	filter := make(map[byte]struct{})
	q := NewQueue()
	for i := 0; i < len(s); i++ {
		b := s[i]
		if _, ok := filter[b]; !ok {
			q.Push(b)
			if q.Length() > maxLength {
				maxLength = q.Length()
			}
			filter[b] = struct{}{}
			continue
		}

		for q.Length() > 0 {
			tmp := q.Pop()
			if tmp == b {
				q.Push(b)
				break
			}
			delete(filter, tmp)
		}
	}
	return maxLength
}

func main() {
	str := "dvdf"
	length := lengthOfLongestSubstring(str)
	fmt.Printf("str:%s,length:%d\n", str, length)
}
