package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 1}
	answer := firstMissingPositive(nums)
	fmt.Println(answer)
}

// 原地哈希
func firstMissingPositive(nums []int) int {
	length := len(nums)
	i := 0
	for i < length {
		if nums[i] > length {
			nums[i] = -1
		}
		if nums[i]-1 == i || nums[i] <= 0 {
			i++
			continue
		}
		if nums[i] == nums[nums[i]-1] {
			nums[i] = -1
		} else {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	for index, num := range nums {
		if index+1 != num {
			return index + 1
		}
	}
	return length + 1
}

//时间复杂度On，空间复杂度On，不符合题意
// func firstMissingPositive(nums []int) int {
// 	maxNum := 0
// 	for _, num := range nums {
// 		if num > 0 && maxNum < num {
// 			maxNum = num
// 		}
// 	}
// 	m := make(map[int]bool, 10000)
// 	for _, num := range nums {
// 		m[num] = true
// 	}
// 	for i := 1; i <= maxNum; i++ {
// 		if !m[i] {
// 			return i
// 		}
// 	}
// 	return maxNum + 1
// }
