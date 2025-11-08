package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	lenPtr := head
	length := 0
	for lenPtr != nil {
		length++
		lenPtr = lenPtr.Next
	}
	ptrHalf := head
	for i := 0; i < length/2; i++ {
		ptrHalf = ptrHalf.Next
	}
	finalPtr := reverseList(ptrHalf)
	for i := 0; i < length/2; i++ {
		if finalPtr.Val != head.Val {
			return false
		}
		finalPtr = finalPtr.Next
		head = head.Next
	}
	return true
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
