package main

import "fmt"

func main() {
	fmt.Println(search([]int{1}, 1))
}
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right-left)/2 + left
		if nums[mid] < target {
			if target > nums[len(nums)-1] && nums[mid] < nums[0] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target <= nums[len(nums)-1] && nums[mid] >= nums[0] && nums[mid] > nums[len(nums)-1] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	if left > len(nums)-1 || left < 0 || nums[left] != target {
		left = -1
	}
	return left
}
