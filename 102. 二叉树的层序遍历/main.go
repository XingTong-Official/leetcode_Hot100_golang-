package main

func main() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	q := NewQueue[*TreeNode]()
	q.Push(root)
	answer := [][]int{}
	if root == nil {
		return answer
	}
	level := 0
	nums1 := 1
	nums2 := 0
	answer = append(answer, []int{})
	for q.length != 0 {
		current := q.Front()
		if current.Left != nil {
			q.Push(current.Left)
			nums2++
		}
		if current.Right != nil {
			q.Push(current.Right)
			nums2++
		}
		answer[level] = append(answer[level], current.Val)
		q.Pop()
		nums1--
		if nums1 == 0 {
			level++
			nums1 = nums2
			nums2 = 0
			answer = append(answer, []int{})
		}
	}
	answer = answer[:len(answer)-1]
	return answer
}

type queue[T any] struct {
	length      int
	queueStruct []T
}

func (q *queue[T]) Push(item T) {
	q.queueStruct = append(q.queueStruct, item)
	q.length++
}
func (q *queue[T]) Pop() {
	q.queueStruct = q.queueStruct[1:]
	q.length--
}
func (q *queue[T]) Front() T {
	return q.queueStruct[0]
}
func NewQueue[T any]() *queue[T] {
	return &queue[T]{
		length:      0,
		queueStruct: make([]T, 0, 16),
	}
}
