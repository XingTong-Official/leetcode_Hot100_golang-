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

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}

// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
// 	var answer *TreeNode
// 	var info1 []*TreeNode
// 	stack := []*TreeNode{}
// 	var LDR func(root *TreeNode)
// 	LDR = func(root *TreeNode) {
// 		if root == nil {
// 			return
// 		}
// 		stack = append(stack, root)
// 		if root.Val == p.Val || root.Val == q.Val {
// 			if len(info1) == 0 {
// 				info1 = append([]*TreeNode{}, stack...)
// 			} else {
// 				var front *TreeNode
// 				l := len(info1)
// 				for i := 0; i < len(stack); i++ {
// 					if i >= l || info1[i] != stack[i] {
// 						answer = front
// 						break
// 					}
// 					if info1[i] == stack[i] {
// 						front = info1[i]
// 					}
// 				}
// 			}
// 		}
// 		LDR(root.Left)
// 		LDR(root.Right)
// 		stack = stack[:len(stack)-1]
// 	}
// 	LDR(root)
// 	return answer
// }
