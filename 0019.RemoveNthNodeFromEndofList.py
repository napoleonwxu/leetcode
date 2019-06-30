# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution(object):
    def removeNthFromEnd(self, head, n):
        """
        :type head: ListNode
        :type n: int
        :rtype: ListNode
        """
        if not head or not n:
            return head
        l = r = head
        for i in xrange(n):
            r = r.next
        if not r:
            return head.next
        while r.next:
            l = l.next
            r = r.next
        l.next = l.next.next
        return head
