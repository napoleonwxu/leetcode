# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def oddEvenList(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        head1 = odd = ListNode(0)
        head2 = even = ListNode(0)
        p = head
        '''
        while p:
            odd.next = p
            odd = odd.next
            p = p.next
            even.next = p
            even = even.next
            if p:
                p = p.next
        if even:
            even.next = None
        '''
        while p:
            odd.next = p
            odd = odd.next
            even.next = p.next
            even = even.next
            p = p.next.next if even else None

        odd.next = head2.next
        return head1.next
