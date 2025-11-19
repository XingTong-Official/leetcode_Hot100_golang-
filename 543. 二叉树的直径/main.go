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

// 基于遍历计算左右树深度，求最大值，效率极差
func diameterOfBinaryTree(root *TreeNode) int {
	m := 0
	var f func(r *TreeNode) int
	f = func(root *TreeNode) int {
		if root == nil {
			return -1
		}
		a := f(root.Left) + 1
		b := f(root.Right) + 1
		m = max(m, a+b)
		return max(a, b)
	}
	f(root)
	return m
}
