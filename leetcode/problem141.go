package main

/**
* @author moorewu
* @E-mail moorewu@tencent.com
* @date 2020/3/6 9:55 上午
 */

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func main() {

}
