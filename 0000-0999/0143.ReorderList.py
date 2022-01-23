# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def reorderList(self, head):
        """
        :type head: ListNode
        :rtype: void Do not return anything, modify head in-place instead.
        """
        if not head or not head.next:
            return
        slow = head
        fast = head.next
        while fast and fast.next:
            slow = slow.next
            fast = fast.next.next
        second = slow.next
        slow.next = None
        p = head
        q = self.reverse(second)
        while p and q:
            tp, tq = p.next, q.next
            p.next = q
            q.next = tp
            p, q = tp, tq

    def reverse(self, head):
        tail = None
        while head:
            temp = head
            head = head.next
            temp.next = tail
            tail = temp
        return tail
