# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution(object):
    def swapPairs(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        if not head or not head.next:
            return head
        '''
        h = ListNode(0)
        h.next = head
        p, q = h, head
        while q and q.next:
            p.next = q.next
            q.next = p.next.next
            p.next.next = q
            p = q
            q = p.next
        return h.next
        '''
        p = head.next
        head.next = self.swapPairs(p.next)
        p.next = head
        return p
