# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def isPalindrome(self, head):
        """
        :type head: ListNode
        :rtype: bool
        """
        if head is None or head.next is None:
            return True
        fast = slow = head
        while fast and fast.next:
            fast = fast.next.next
            slow = slow.next
        sec = None
        while slow:
            p = slow
            slow = slow.next
            p.next = sec
            sec = p
        while head and sec:
            if head.val != sec.val:
                return False
            head = head.next
            sec = sec.next
        return True
        