# Definition for singly-linked list with a random pointer.
# class RandomListNode(object):
#     def __init__(self, x):
#         self.label = x
#         self.next = None
#         self.random = None

class Solution(object):
    def copyRandomList(self, head):
        """
        :type head: RandomListNode
        :rtype: RandomListNode
        """
        if not head:
            return
        p = head
        while p:
            pp = RandomListNode(p.label)
            pp.next = p.next
            p.next = pp
            p = pp.next

        p = head
        while p:
            pp = p.next
            if p.random:
                pp.random = p.random.next
            p = pp.next

        headNew = head.next
        p, pp = head, headNew
        while pp.next:
            p.next = pp.next
            p = p.next
            pp.next = p.next
            pp = pp.next
        p.next = None
        #pp.next = None

        return headNew
