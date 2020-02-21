package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	cur := head
	overflow := 0
	for l1 != nil || l2 != nil || overflow != 0 {
		v1, v2 := 0, 0
		if l1 != nil {
			v1, l1 = v1+l1.Val, l1.Next
		}
		if l2 != nil {
			v2, l2 = v2+l2.Val, l2.Next
		}
		n := v1 + v2 + overflow
		if n >= 10 {
			cur.Next = &ListNode{Val: n - 10, Next: nil}
			overflow = 1
		} else {
			cur.Next = &ListNode{Val: n, Next: nil}
			overflow = 0
		}
		cur = cur.Next
	}
	return head.Next
}
//
//func main() {
//
//}
