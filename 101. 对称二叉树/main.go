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

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	reserveTree := invertTree(root.Left)
	return isEqualTree(reserveTree, root.Right)
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
func isEqualTree(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if (root1 == nil && root2 != nil) || (root1 != nil && root2 == nil) {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}
	return isEqualTree(root1.Left, root2.Left) && isEqualTree(root1.Right, root2.Right)

}
