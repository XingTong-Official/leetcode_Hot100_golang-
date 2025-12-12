package main

import "fmt"

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 100}, 8))
}
func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left <= right {
		middle := (right-left)/2 + left
		if nums[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	var first int
	if left >= 0 && left <= len(nums)-1 && nums[left] == target {
		first = left
	} else {
		first = -1
	}

	left, right = 0, len(nums)-1
	for left <= right {
		middle := (right-left)/2 + left
		if nums[middle] > target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	if right < 0 || right > len(nums)-1 || nums[right] != target {
		right = -1
	}
	return []int{first, right}
}
