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

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	answer := head.Next
	head.Next = head.Next.Next
	answer.Next = head
	for head.Next != nil && head.Next.Next != nil {
		temp := head.Next
		head.Next = head.Next.Next
		temp.Next = head.Next.Next
		head.Next.Next = temp
		head = head.Next
		head = head.Next
	}
	return answer
}
