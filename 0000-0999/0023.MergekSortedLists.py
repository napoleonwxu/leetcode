# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution(object):
    def mergeKLists(self, lists):
        """
        :type lists: List[ListNode]
        :rtype: ListNode
        """
        k = len(lists)
        #return self.depart(lists, 0, k-1)
        if k == 0:
            return lists
        while k>1:
            for i in xrange(k/2):
                lists[i] = self.mergeTwoLists(lists[i], lists[k-i-1])
            k = (k+1)/2
        return lists[0]

    def depart(self, lists, start, end):
        if start == end:
            return lists[start]
        if start < end:
            mid = (start+end)/2
            l1 = self.depart(lists, start, mid)
            l2 = self.depart(lists, mid+1, end)
            return self.mergeTwoLists(l1, l2)

    def mergeTwoLists(self, l1, l2):
        p = head = ListNode(0)
        while l1 and l2:
            if l1.val <= l2.val:
                p.next = l1
                l1 = l1.next
            else:
                p.next = l2
                l2 = l2.next
            p = p.next
        if l1:
            p.next = l1
        if l2:
            p.next = l2
        return head.next