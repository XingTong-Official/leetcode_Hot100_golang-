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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fastPtr, slowptr := head, head
	for i := 0; i < n; i++ {
		fastPtr = fastPtr.Next
	}
	if fastPtr == nil {
		return head.Next
	}
	for fastPtr.Next != nil {
		fastPtr = fastPtr.Next
		slowptr = slowptr.Next
	}
	slowptr.Next = slowptr.Next.Next
	return head
}
