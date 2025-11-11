package main

func main() {

}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	m1 := make(map[int]*Node, 500)
	temp := head
	m2 := make(map[*Node]int, 500)
	index := 0
	for temp != nil {
		m2[temp] = index
		temp = temp.Next
		m1[index] = &Node{}
		index++
	}
	index = 0
	for head != nil {
		m1[index].Val = head.Val
		m1[index].Next = m1[index+1]
		if head.Random == nil {
			m1[index].Random = nil
		} else {
			m1[index].Random = m1[m2[head.Random]]
		}

		head = head.Next
		index++
	}
	return m1[0]
}
