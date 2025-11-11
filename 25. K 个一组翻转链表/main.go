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

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	answer := head
	for i := 0; i < k-1; i++ {
		answer = answer.Next
	}
	final := &ListNode{}
	ptr := head
	for {
		//检验是否为最后的链表部分
		ptrLength := ptr
		for i := 0; i < k-1; i++ {
			if ptrLength == nil || ptrLength.Next == nil {
				final.Next = ptr
				return answer
			}
			ptrLength = ptrLength.Next
		}
		temp := ptrLength.Next
		ptrLength.Next = nil
		final.Next = reverseList(ptr)
		final = ptr
		ptr = temp
	}

}
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	answer := reverseListD(head)
	head.Next = nil
	return answer
}
func reverseListD(head *ListNode) *ListNode {
	if head.Next.Next == nil {
		head.Next.Next = head
		return head.Next
	} else {
		temp := reverseListD(head.Next)
		head.Next.Next = head
		return temp
	}
}
