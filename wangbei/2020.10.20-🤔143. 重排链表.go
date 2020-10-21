/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	middle := FindMiddle(head)
	head2 := middle.Next
	middle.Next = nil
	reverseHead2 := Reverse(head2)
	head = MergeTwoList(head, reverseHead2)
}
func FindMiddle(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	fast := head.Next
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
func Reverse(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		temp := head.Next
		head.Next = pre
		pre = head
		head = temp
	}
	return pre
}
func MergeTwoList(head1, head2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	cur := dummy
	for head1 != nil && head2 != nil {
		cur.Next = head1
		cur = cur.Next
		head1 = head1.Next

		cur.Next = head2
		cur = cur.Next
		head2 = head2.Next
	}
	if head1 != nil {
		cur.Next = head1
		cur = cur.Next
		head1 = head1.Next
	}
	if head2 != nil {
		cur.Next = head2
		cur = cur.Next
		head2 = head2.Next
	}
	return dummy.Next
}