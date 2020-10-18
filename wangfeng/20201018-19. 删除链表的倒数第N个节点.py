# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def removeNthFromEnd(self, head: ListNode, n: int) -> ListNode:
        dummy = ListNode(val=-1, next=head)
        dele = dummy
        tail = dummy
        for i in range(n):
            tail = tail.next
        while tail.next != None:
            dele = dele.next
            tail =  tail.next
        dele.next = dele.next.next
        return dummy.next
