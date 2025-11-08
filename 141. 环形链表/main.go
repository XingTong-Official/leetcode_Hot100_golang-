package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	fastPtr := head.Next.Next
	slowPtr := head
	for slowPtr != fastPtr {
		if fastPtr == nil || fastPtr.Next == nil {
			return false
		}
		fastPtr = fastPtr.Next.Next
		slowPtr = slowPtr.Next
	}
	return true
}
