# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def rotateRight(self, head, k):
        """
        :type head: ListNode
        :type k: int
        :rtype: ListNode
        """
        if not head:
            return head
        l = 1
        tail = head
        while tail.next:
            l += 1
            tail = tail.next
        tail.next = head
        k %= l
        for i in xrange(l-k):
            tail = tail.next
        head = tail.next
        tail.next = None
        return head
        