package main


func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	cur := head
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				cur.Next = &ListNode{l1.Val, nil}
				l1 = l1.Next
			} else {
				cur.Next = &ListNode{l2.Val, nil}
				l2 = l2.Next
			}
		} else {
			if l1 != nil {
				cur.Next = &ListNode{l1.Val, nil}
				l1 = l1.Next
			} else {
				cur.Next = &ListNode{l2.Val, nil}
				l2 = l2.Next
			}
		}

		cur = cur.Next
	}
	return head.Next
}

