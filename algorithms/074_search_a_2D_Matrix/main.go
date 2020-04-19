package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	um := 0
	dm := len(matrix) - 1
	if target < matrix[um][0] || target > matrix[dm][len(matrix[dm])-1] {
		return false
	}
	for um <= dm {
		if matrix[um][0] == target || matrix[dm][0] == target {
			return true
		}
		mid := (um + dm) / 2
		if target > matrix[mid][0] {
			if target <= matrix[mid][len(matrix[mid])-1] {
				um = mid
				break
			}
			um = mid + 1
		} else if target < matrix[mid][0] {
			dm = mid - 1
		} else {
			return true
		}
	}

	lindex := 0
	rindex := len(matrix[um]) - 1
	if target < matrix[um][lindex] || target > matrix[um][rindex] {
		return false
	}
	for lindex <= rindex {
		if matrix[um][lindex] == target || matrix[um][rindex] == target {
			return true
		}
		mid := (lindex + rindex) / 2
		if target < matrix[um][mid] {
			rindex = mid - 1
		} else if target > matrix[um][mid] {
			lindex = mid + 1
		} else {
			return true
		}
	}
	return false
}

func main() {
	matrix := [][]int{[]int{1}}
	target := 1
	exist := searchMatrix(matrix, target)
	fmt.Printf("matrix:%+v,target:%d,exist:%+v\n", matrix, target, exist)
}
