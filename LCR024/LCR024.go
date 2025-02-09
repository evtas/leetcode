package LCR024

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

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	headNode := new(ListNode)
	headNode.Next = head
	cur := headNode.Next
	headNode.Next = nil
	for cur != nil {
		pre := cur
		cur = cur.Next

		pre.Next = headNode.Next
		headNode.Next = pre
	}
	return headNode.Next
}
