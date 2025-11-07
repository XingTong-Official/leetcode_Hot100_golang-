package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4}
	answer := productExceptSelf(nums)
	fmt.Println(answer)
}
func productExceptSelf(nums []int) []int {
	length := len(nums)
	answer := make([]int, length)

	pre := make([]int, length)
	pre[0] = 1
	for index := 1; index <= length-1; index++ {
		pre[index] = pre[index-1] * nums[index-1]
	}

	suf := make([]int, length)
	suf[length-1] = 1
	for index := length - 2; index >= 0; index-- {
		suf[index] = suf[index+1] * nums[index+1]
	}
	for index := 0; index <= length-1; index++ {
		answer[index] = pre[index] * suf[index]
	}
	return answer
}

//[-1,1,0,-3,3]
//[3,-3,0,1,-1]
