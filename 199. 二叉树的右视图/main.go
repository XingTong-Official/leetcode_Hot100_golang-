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

func rightSideView(root *TreeNode) []int {
	answer := []int{}
	LO := LevelOrder(root)
	for _, nums := range LO {
		answer = append(answer, nums[len(nums)-1])
	}
	return answer
}
func LevelOrder(root *TreeNode) [][]int {
	answer := [][]int{}
	if root == nil {
		return answer
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	ptr := 1
	next := 0
	for len(queue) != 0 {

		temp := []int{}
		for i := 0; i < ptr; i++ {
			current := queue[0]
			if current.Left != nil {
				queue = append(queue, current.Left)
				next++
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
				next++
			}
			temp = append(temp, current.Val)
			queue = queue[1:]
		}
		ptr = next
		next = 0
		answer = append(answer, temp)
	}
	return answer
}
