# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def addTwoNumbers(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
        s1, s2 = self.toNum(l1), self.toNum(l2)
        c = 0
        head = None
        while s1 or s2 or c:
            if s1:
                c += s1.pop()
            if s2:
                c += s2.pop()
            tmp = ListNode(c%10)
            tmp.next = head
            head = tmp
            c /= 10
        return head
        
    def toNum(self, l):
        stack = []
        while l:
            stack.append(l.val)
            l = l.next
        return stack