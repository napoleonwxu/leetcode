class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution(object):
    def addListSum(self, l1, l2):
        l = self.addList(self.reverse(l1), self.reverse(l2))
        ans = self.reverse(l)
        return ans

    def reverse(self, l):
        head = None
        while l:
            t = l
            l = l.next
            t.next = head
            head = t
        return head

    def addList(self, l1, l2):
        head = p = ListNode(0)
        c = 0
        while l1 or l2 or c:
            if l1:
                c += l1.val
                l1 = l1.next
            if l2:
                c += l2.val
                l2 = l2.next
            p.next = ListNode(c%10)
            p = p.next
            c /= 10
        return head.next

    def printList(self, head):
        while head:
            print head.val,
            head = head.next
        print

l1 = ListNode(4)
l1.next = ListNode(2)
l2 = ListNode(5)
l2.next = ListNode(8)
test = Solution()
add1 = test.addList(l1, l2)
add2 = test.addListSum(l1, l2)
test.printList(add1)
test.printList(add2)
