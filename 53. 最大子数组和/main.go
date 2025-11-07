package main

import (
	"fmt"
)

func main() {
	nums := []int{-2, -1}
	answer := maxSubArray(nums)
	fmt.Println(answer)
}
func maxSubArray(nums []int) int {
	maxValue := nums[0]
	q := queueStruct{}
	for _, num := range nums {
		q.push(num)
		if q.value > maxValue {
			maxValue = q.value
		}
		if q.value <= 0 {
			q.clear()
		}

	}
	return maxValue
}

type queueStruct struct {
	Queue []int
	value int
}

func (q *queueStruct) clear() {
	q.Queue = []int{}
	q.value = 0
}
func (q *queueStruct) push(num int) {
	q.Queue = append(q.Queue, num)
	q.value += num
}
func (q *queueStruct) pop() {
	q.value -= q.Queue[0]
	q.Queue = q.Queue[1:]
}
func (q *queueStruct) top() int {
	return q.Queue[0]
}
