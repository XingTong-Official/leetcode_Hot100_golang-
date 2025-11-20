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

func kthSmallest(root *TreeNode, k int) int {
	nums := []int{}
	var LDR func(root *TreeNode)
	LDR = func(root *TreeNode) {
		if root == nil {
			return
		}
		LDR(root.Left)
		nums = append(nums, root.Val)
		LDR(root.Right)
	}
	LDR(root)
	return nums[k-1]
}
