/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reverseKGroup(head *ListNode, k int) *ListNode {
    if head == nil || k <= 1 {
        return head
    }
    p := head
    for i := 0; i < k; i++ {
        if p == nil {
            return head
        }
        p = p.Next
    }
    newHead := reverseKGroup(p, k)
    for i := 0; i < k; i++ {
        nxt := head.Next
        head.Next = newHead
        newHead = head
        head = nxt
    }
    return newHead
}