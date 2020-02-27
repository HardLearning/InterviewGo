package main

/**
* @author moorewu
* @E-mail moorewu@tencent.com
* @date 2020/2/27 8:55 上午
 */

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	prev := new(ListNode)
	cur := head
	head = cur.Next
	for ; cur != nil && cur.Next != nil; cur = cur.Next {
		next := cur.Next
		if prev != nil {
			prev.Next = next
		}
		cur.Next, next.Next, prev = next.Next, cur, cur
	}
	return head
}
