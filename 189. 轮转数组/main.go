package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)
	fmt.Println(nums)
}
func rotate(nums []int, k int) {
	k = k % len(nums)
	tmp := make([]int, len(nums)-k)
	copy(tmp, nums[:len(nums)-k])
	copy(nums[:k], nums[len(nums)-k:])
	copy(nums[k:], tmp)
}
