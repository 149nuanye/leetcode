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

func minWindow(s string, t string) string {
	tarHash := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		if num, ok := tarHash[t[i]]; ok {
			tarHash[t[i]] = num + 1
			continue
		}
		tarHash[t[i]] = 1
	}
	lindex := -1
	minLength := len(s)
	queue := NewQueue()
	tmpHash := make(map[byte]int)
	for j := 0; j < len(s); j++ {
		if _, ok := tarHash[s[j]]; !ok {
			queue.Push(s[j])
			continue
		}
		// fmt.Printf("1: s[j]:%+v,j:%d,tmpHash:%+v\n", s[j], j, tmpHash)
		tmpHash[s[j]] = tmpHash[s[j]] + 1
		queue.Push(s[j])
		full := true
		for key := range tarHash {
			if tmpHash[key] < tarHash[key] {
				full = false
			}
		}
		if !full {
			continue
		}
		// fmt.Printf("2: tarHash:%+v\n", tarHash)
		if queue.Length() <= minLength {
			minLength = queue.Length()
			lindex = j - queue.Length() + 1
		}
		// fmt.Printf("3: minLength:%+v,lindex:%d\n", minLength, lindex)
		for queue.Length() > 0 {
			value := queue.Pop()
			if num, ok := tmpHash[value]; ok {
				tmpHash[value] = num - 1
				if num-1 < tarHash[value] {
					if queue.Length()+1 < minLength {
						minLength = queue.Length() + 1
						lindex = j - queue.Length()
					}
					break
				}
			}
		}
	}
	if lindex < 0 {
		return ""
	}
	return s[lindex : lindex+minLength]
}

func main() {
	s := "AC"
	t := "ABC"
	minStr := minWindow(s, t)
	fmt.Printf("s:%s,t:%s,min:%s\n", s, t, minStr)
}
