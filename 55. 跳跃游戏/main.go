package main

import "fmt"

func main() {
	fmt.Println(canJump([]int{1, 0, 2}))
}
func canJump(nums []int) bool {
	mx := 0
	for i, num := range nums {
		if mx < i {
			return false
		}
		mx = max(i+num, mx)
	}
	return true
}

// func canJump(nums []int) bool {
// 	local := 0
// 	for local < len(nums)-1 {
// 		num := nums[local]
// 		if num == 0 {
// 			return false
// 		}
// 		if num+local >= len(nums)-1 {
// 			return true
// 		}
// 		m := 0
// 		temp := local
// 		for i := 1; i <= num; i++ {
// 			if nums[local+i]+local+i >= m {
// 				m = nums[local+i] + local + i
// 				temp = local + i
// 			}
// 		}
// 		local = temp
// 	}
// 	return true
// }
