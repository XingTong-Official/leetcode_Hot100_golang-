package main

import (
	"fmt"
)

func main() {
	nums := []int{1, -1, 0}
	k := 0
	answer := subarraySum(nums, k)
	fmt.Println(answer)
}

// 前缀和
func subarraySum(nums []int, k int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	answer, pre := 0, 0
	m := make(map[int]int, length+1)
	m[0] = 1
	for i := 0; i < length; i++ {
		pre += nums[i]
		mnum, ok := m[pre-k]
		if ok {
			answer += mnum
		}
		m[pre]++
	}
	return answer
}

//遍历循环,n²
// func subarraySum(nums []int, k int) int {
// 	answer := 0
// 	for left, _ := range nums {
// 		total := 0
// 		for right := left; right < len(nums); right++ {
// 			total += nums[right]
// 			if total == k {
// 				answer++
// 			}
// 		}
// 	}
// 	return answer
// }

// func calculate(nums []int) int {
// 	total := 0
// 	for _, num := range nums {
// 		total += num
// 	}
// 	return total
// }
