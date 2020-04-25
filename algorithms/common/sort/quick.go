package main

func quickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	mid := nums[0]
	lindex := 0
	rindex := len(nums) - 1
	for lindex < rindex {
		for rindex > lindex {
			if nums[rindex] < mid {
				nums[lindex] = nums[rindex]
				lindex++
				break
			}
			rindex--
		}
		for rindex > lindex {
			if nums[lindex] >= mid {
				nums[rindex] = nums[lindex]
				rindex--
				break
			}
			lindex++
		}
	}
	nums[lindex] = mid
	quickSort(nums[:lindex])
	if lindex < len(nums)-1 {
		quickSort(nums[lindex+1:])
	}
}
