package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	ptr1, ptr2 := list1, list2
	answer := &ListNode{}
	if ptr1 == nil {
		return ptr2
	}
	if ptr2 == nil {
		return ptr1
	}
	if ptr1.Val > ptr2.Val {
		answer = ptr2
		ptr2 = ptr2.Next
	} else {
		answer = ptr1
		ptr1 = ptr1.Next
	}
	answer1 := answer
	for {
		if ptr1 == nil {
			answer.Next = ptr2
			break
		}
		if ptr2 == nil {
			answer.Next = ptr1
			break
		}
		if ptr1.Val > ptr2.Val {
			answer.Next = ptr2
			ptr2 = ptr2.Next
			answer = answer.Next
		} else {
			answer.Next = ptr1
			ptr1 = ptr1.Next
			answer = answer.Next
		}
	}
	return answer1
}
