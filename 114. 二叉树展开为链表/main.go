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

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	current := root
	for current.Right != nil || current.Left != nil {
		if current.Left != nil {
			temp := current.Left
			for temp.Right != nil {
				temp = temp.Right
			}
			current.Right, temp.Right = current.Left, current.Right
			current.Left = nil
		}
		current = current.Right
	}
}

// func flatten(root *TreeNode) {
// 	if root == nil {
// 		return
// 	}
// 	nums := []*TreeNode{}
// 	var DLR func(root *TreeNode)

// 	DLR = func(root *TreeNode) {
// 		if root == nil {
// 			return
// 		}
// 		nums = append(nums, root)
// 		DLR(root.Left)
// 		DLR(root.Right)
// 	}
// 	DLR(root)
// 	length := len(nums)
// 	for i := 0; i < length-1; i++ {
// 		nums[i].Right = nums[i+1]
// 		nums[i].Left = nil
// 	}
// 	nums[length-1].Left = nil
// 	nums[length-1].Right = nil
// }
