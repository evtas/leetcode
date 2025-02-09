package _9

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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	canMove := false
	count := 0
	tail := head
	prev := head
	for tail != nil {
		if canMove {
			prev = prev.Next
		}
		tail = tail.Next

		count += 1
		if count == n+1 {
			canMove = true
		}
	}
	if count == n {
		return head.Next
	}
	if prev.Next == nil {
		prev = nil
	}
	prev.Next = prev.Next.Next
	return head
}
