package main

import "fmt"

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	answer := longestConsecutive(nums)
	fmt.Println(answer)
}

/*
设计一个map，存放所有数，核心思想是，寻找每个队列的起始节点，如果不是起始节点，就跳过，对于起始节点
判断其长度
**/
func longestConsecutive(nums []int) int {
	set := make(map[int]bool, len(nums))
	for _, i := range nums {
		set[i] = true
	}
	ans := 0
	for i := range set {
		if set[i-1] {
			continue
		}
		y := i + 1
		for set[y] {
			y++
		}
		ans = max(ans, y-i)
	}
	return ans
}
