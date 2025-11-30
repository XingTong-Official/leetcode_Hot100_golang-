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

func maxPathSum(root *TreeNode) int {
	m := -1001
	var LDR func(root *TreeNode) int
	LDR = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := LDR(root.Left)
		right := LDR(root.Right)
		if right <= 0 && left <= 0 {
			m = max(m, root.Val)
			return root.Val
		} else {
			m = max(m, root.Val+left, root.Val+right, left+right+root.Val)
			return max(root.Val+left, root.Val+right)
		}
	}
	LDR(root)
	return m
}
