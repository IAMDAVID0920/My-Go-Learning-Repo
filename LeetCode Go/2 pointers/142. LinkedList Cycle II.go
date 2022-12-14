package main

/*
*
  - Definition for singly-linked list.

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			fast = head
			for slow != fast {
				fast = fast.Next
				slow = slow.Next
			}
			return slow
		}
	}
	return nil
}
