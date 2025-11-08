package main

func main() {

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	add := 0
	list := &ListNode{}
	answer := list
	for {
		if l1 == nil && l2 == nil && add == 0 {
			break
		}
		value := 0
		value, add = sum(l1, l2, add)
		newPoint := &ListNode{
			Val:  value,
			Next: nil,
		}
		list.Next = newPoint
		list = list.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	return answer.Next
}
func sum(i, j *ListNode, addNum int) (value int, add int) {
	i1, j1 := 0, 0
	if i == nil {
		i1 = 0
	} else {
		i1 = i.Val
	}
	if j == nil {
		j1 = 0
	} else {
		j1 = j.Val
	}
	return (addNum + i1 + j1) % 10, (addNum + i1 + j1) / 10
}
