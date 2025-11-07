package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	//求出指针位置差值
	lengthA := 0
	ptrA := headA
	for {
		if ptrA.Next == nil {
			break
		}
		lengthA++
		ptrA = ptrA.Next
	}
	lengthB := 0
	ptrB := headB
	for {
		if ptrB.Next == nil {
			break
		}
		lengthB++
		ptrB = ptrB.Next
	}

	//消除指针位置差异
	if lengthA > lengthB {
		for i := 0; i < lengthA-lengthB; i++ {
			headA = headA.Next
		}
	} else if lengthA < lengthB {
		for i := 0; i < lengthB-lengthA; i++ {
			headB = headB.Next
		}
	}
	for headA != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}
	return nil
}
