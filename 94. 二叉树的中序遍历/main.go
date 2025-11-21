package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type stack struct {
	arrs   []*TreeNode
	length int
}

func NewStack() *stack {
	return &stack{arrs: make([]*TreeNode, 0, 40), length: 0}
}
func (s *stack) Pop() *TreeNode {
	value := s.Top()
	s.arrs = s.arrs[:s.length-1]
	s.length--
	return value
}
func (s *stack) Top() *TreeNode {
	return s.arrs[s.length-1]
}
func (s *stack) Push(node *TreeNode) {
	s.arrs = append(s.arrs, node)
	s.length++
}

// 迭代
func inorderTraversal(root *TreeNode) []int {
	stack := NewStack()
	answer := []int{}
	for root != nil || stack.length != 0 {
		for root != nil {
			stack.Push(root)
			root = root.Left
		}
		root = stack.Pop()
		answer = append(answer, root.Val)
		root = root.Right
	}
	return answer
}

//先序遍历
// func inorderTraversal(root *TreeNode) []int {
// 	answer := []int{}
// 	s := NewStack()
// 	s.Push(root)
// 	for s.length != 0 {
// 		current := s.Pop()
// 		answer = append(answer, current.Val)
// 		if current.Right != nil {
// 			s.Push(current.Right)
// 		}
// 		if current.Left != nil {
// 			s.Push(current.Left)
// 		}
// 	}
// 	return answer
// }

//闭包的递归
// func inorderTraversal(root *TreeNode) []int {
// 	answer := []int{}
// 	var traverse func(*TreeNode)
// 	traverse = func(node *TreeNode) {
// 		if node == nil {
// 			return
// 		}
// 		traverse(node.Left)
// 		answer = append(answer, node.Val)
// 		traverse(node.Right)
// 	}
// 	traverse(root)
// 	return answer
// }
