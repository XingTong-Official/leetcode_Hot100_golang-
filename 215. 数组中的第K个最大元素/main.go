package main

import "container/heap"

func main() {

}

type PriorityQueue []int

func (q PriorityQueue) Len() int {
	return len(q)
}
func (q PriorityQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}
func (q PriorityQueue) Less(i, j int) bool {
	return q[i] > q[j]
}
func (q *PriorityQueue) Push(x any) {
	i := x.(int)
	*q = append(*q, i)
}
func (q *PriorityQueue) Pop() any {
	n := q.Len()
	num := (*q)[n-1]
	*q = (*q)[:n-1]
	return num
}
func findKthLargest(nums []int, k int) int {
	a := PriorityQueue(nums)
	heap.Init(&a)
	for i := 1; i < k; i++ {
		heap.Pop(&a)
	}
	return heap.Pop(&a).(int)
}
