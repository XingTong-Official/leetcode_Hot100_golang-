package main

func main() {
	searchInsert([]int{1, 3, 5, 6}, 2)
}
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for {
		local := (right-left)/2 + left
		if left >= right {
			if nums[local] < target {
				return local + 1
			}
			return local
		}
		if nums[local] == target {
			return local
		}

		if nums[local] > target {
			right = (right-left)/2 + left
		} else if nums[local] < target {
			left = (right-left)/2 + left + 1
		}
	}
}
