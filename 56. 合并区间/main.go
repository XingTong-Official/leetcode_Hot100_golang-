package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := [][]int{{1, 4}, {0, 2}, {3, 5}}
	answer := merge(nums)
	fmt.Println(answer)
}
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	slices.SortFunc(intervals, func(a, b []int) int { return a[0] - b[0] })
	answer := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= answer[len(answer)-1][1] {
			answer[len(answer)-1] = []int{answer[len(answer)-1][0], max(intervals[i][1], answer[len(answer)-1][1])}
		} else {
			answer = append(answer, intervals[i])
		}
	}
	return answer
}
