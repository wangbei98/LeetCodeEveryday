/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: 0}
	dummy.Next = head
	p, q := head, head
	pre := dummy
	for i := 0; i < n; i++ {
		p = p.Next
	}
	for p != nil {
		p = p.Next
		pre = q
		q = q.Next
	}
	pre.Next = q.Next
	return dummy.Next
}