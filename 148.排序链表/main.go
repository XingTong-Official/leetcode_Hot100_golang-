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

func sortList(head *ListNode) *ListNode {
	//加置空头结点
	nummy := &ListNode{
		Val:  0,
		Next: head,
	}
	ptr := nummy
	return sort(ptr).Next
}
func sort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{Val: 0, Next: head}

	// 获取链表长度
	length := 0
	for curr := head; curr != nil; curr = curr.Next {
		length++
	}

	// 冒泡排序
	for i := 0; i < length; i++ {
		prev := dummy
		curr := dummy.Next

		for j := 0; j < length-i-1 && curr.Next != nil; j++ {
			if curr.Val > curr.Next.Val {
				// 交换节点
				next := curr.Next
				curr.Next = next.Next
				next.Next = curr
				prev.Next = next

				prev = next
			} else {
				prev = curr
				curr = curr.Next
			}
		}
	}

	return dummy.Next
}
