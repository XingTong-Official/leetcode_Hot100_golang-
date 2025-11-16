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

func diameterOfBinaryTree(root *TreeNode) int {

}

//基于遍历计算左右树深度，求最大值，效率极差
// func diameterOfBinaryTree(root *TreeNode) int {
// 	max := 0
// 	var f func(r *TreeNode)
// 	f = func(root *TreeNode) {
// 		if root == nil {
// 			return
// 		}
// 		a := maxDepth(root.Left)
// 		b := maxDepth(root.Right)
// 		if a+b > max {
// 			max = a + b
// 		}
// 		f(root.Left)
// 		f(root.Right)
// 	}
// 	f(root)
// 	return max
// }

// type stack struct {
// 	arrs   []*TreeNode
// 	length int
// }

// func NewStack() *stack {
// 	return &stack{arrs: make([]*TreeNode, 0, 40), length: 0}
// }
// func (s *stack) Pop() *TreeNode {
// 	value := s.Top()
// 	s.arrs = s.arrs[:s.length-1]
// 	s.length--
// 	return value
// }
// func (s *stack) Top() *TreeNode {
// 	return s.arrs[s.length-1]
// }
// func (s *stack) Push(node *TreeNode) {
// 	s.arrs = append(s.arrs, node)
// 	s.length++
// }

// func maxDepth(root *TreeNode) int {
// 	s := NewStack()
// 	var D func(root *TreeNode)
// 	max := 0
// 	D = func(root *TreeNode) {
// 		if root == nil {
// 			return
// 		}
// 		s.Push(root)
// 		if s.length > max {
// 			max = s.length
// 		}
// 		D(root.Left)
// 		D(root.Right)
// 		s.Pop()
// 	}
// 	D(root)
// 	return max
// }
