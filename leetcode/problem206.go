package main

/**
* @author moorewu
* @E-mail moorewu@tencent.com
* @date 2020/3/6 11:37 上午
 */

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode
	cur := head
	for cur != nil {
		cur.Next, prev, cur = prev, cur, cur.Next
	}
	return prev
}

func main() {

}
