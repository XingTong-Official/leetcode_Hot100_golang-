package main

import "fmt"

func main() {
	fmt.Println(findMin([]int{5, 1, 2, 3, 4}))
}
func findMin(nums []int) int {
	min := nums[0]
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right-left)/2 + left
		if nums[mid] > min {
			if nums[mid] >= nums[0] && nums[mid] > nums[len(nums)-1] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if nums[mid] >= nums[0] && nums[mid] > nums[len(nums)-1] {
				left = mid + 1
			} else {
				right = mid - 1
			}
			min = nums[mid]
		}
	}
	return min
}
