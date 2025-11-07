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
		temp := reverseList(head.Next)
		head.Next.Next = head
		return temp
	}
}
