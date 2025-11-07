package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	answer := trap(nums)
	fmt.Println(answer)
}

//单调栈思路，第一层循环确保遍历每个元素，第二个循环遍历栈内比当前元素小的值
func trap(height []int) int {
	answer := 0
	stack := []int{}
	for i, h := range height {
		for len(stack) > 0 && h > height[stack[len(stack)-1]] {
			//当无左遮挡时，POP，然后结束此循环
			if len(stack)-2 < 0 {
				stack = stack[:len(stack)-1]
				break
			}
			//存在左遮挡，计算左遮挡Left，top是当前的待计算的深度
			top := len(stack) - 1
			left := top - 1
			//核心计算公式
			ans := (min(h, height[stack[left]]) - height[stack[top]]) * (i - stack[left] - 1)
			answer += ans
			//栈顶出栈
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return answer
}
