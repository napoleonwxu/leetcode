/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reverseBetween(head *ListNode, m int, n int) *ListNode {
    if head == nil || m == n {
        return head
    }
    dummy := &ListNode{}
    dummy.Next = head
    p1 := dummy
    for i := 1; i < m; i++ {
        p1 = p1.Next
    }
    p2 := p1.Next
    p := p2
    for i := m; i < n; i++ {
        p = p.Next
    }
    p3 := p.Next
    p.Next = nil
    for p2 != nil {
        nxt := p2.Next
        p2.Next = p3
        p3 = p2
        p2 = nxt
    }
    p1.Next = p3
    return dummy.Next
}