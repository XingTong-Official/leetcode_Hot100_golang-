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

func isValidBST(root *TreeNode) bool {
	midNums := []int{}
	var MID func(root *TreeNode)
	MID = func(root *TreeNode) {
		if root == nil {
			return
		}
		MID(root.Left)
		midNums = append(midNums, root.Val)
		MID(root.Right)
	}
	MID(root)
	front := midNums[0]
	for i := 1; i < len(midNums); i++ {
		if midNums[i] <= front {
			return false
		}
		front = midNums[i]
	}
	return true
}

// func isValidBST(root *TreeNode) bool {
// 	answer := true
// 	var BST func(root *TreeNode) (min, max int)
// 	BST = func(root *TreeNode) (min, max int) {
// 		if root.Left == nil && root.Right == nil {
// 			return root.Val, root.Val
// 		}
// 		leftVal, rightVal := 0, 0
// 		if root.Left == nil {
// 			rightVal = BST(root.Right)
// 			if rightVal <= root.Val {
// 				answer = false
// 			}
// 			return max(rightVal, root.Val)
// 		}
// 		leftVal = BST(root.Left)
// 		if leftVal >= root.Val || root.Left.Val >= root.Val {
// 			answer = false
// 		}
// 		if root.Right == nil {
// 			leftVal = BST(root.Left)
// 			if leftVal >= root.Val {
// 				answer = false
// 			}
// 			return max(leftVal, root.Val)
// 		}
// 		rightVal = BST(root.Right)
// 		if rightVal <= root.Val || root.Right.Val <= root.Val {
// 			answer = false
// 		}
// 		return max(leftVal, rightVal, root.Val)
// 	}
// 	BST(root)
// 	return answer
// }
