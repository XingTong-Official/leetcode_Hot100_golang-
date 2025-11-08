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

func detectCycle(head *ListNode) *ListNode {

}

// 方法二，设入环前的长度为a，使用快慢指针fast和slow，快指针为慢指针速度两倍，入环点到相遇点
// 距离为b，环的剩下长度为c，可知，2a+2b=a+n*(b+c)+b
// 公式可得到：a=(n-1)b+nc
// a=c+(n-1)(b+c)
// 故有：设指针ptr为链表起始，ptr与slow同时走相同长度，当ptr与slow相遇时，此时相遇点即为入环口
// 此种方法，时间复杂度为On，空间复杂度为O1
func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
		if fast == nil {
			return nil
		}
		fast = fast.Next
		if fast == slow {
			for {
				if head == slow {
					return head
				}
				head = head.Next
				slow = slow.Next
			}
		}
	}
	return nil
}

//方法1：空间复杂度和时间复杂度均为on,map要预留一定空间大小，避免频繁扩容
// func detectCycle(head *ListNode) *ListNode {
// 	m := make(map[*ListNode]bool, 400)
// 	for head != nil {
// 		if m[head] == true {
// 			return head
// 		} else {
// 			m[head] = true
// 		}
// 		head = head.Next
// 	}
// 	return nil
// }
