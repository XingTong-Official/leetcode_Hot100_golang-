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

func invertTree(root *TreeNode) *TreeNode {
	var Reserve func(root *TreeNode)
	Reserve = func(root *TreeNode) {
		if root == nil {
			return
		}
		Reserve(root.Left)
		Reserve(root.Right)
		root.Left, root.Right = root.Right, root.Left
	}
	Reserve(root)
	return root
}
