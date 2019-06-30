# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def reverseKGroup(self, head, k):
        """
        :type head: ListNode
        :type k: int
        :rtype: ListNode
        """
        if not head or k <= 1:
            return head
        p = head
        for i in xrange(k):
            if not p:
                return head
            p = p.next
        pre = self.reverseKGroup(p, k)
        nxt = head.next
        for i in xrange(k-1):
            head.next = pre
            pre = head
            head = nxt
            nxt = head.next
        head.next = pre
        return head
