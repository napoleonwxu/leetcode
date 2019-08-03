/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    head := &ListNode{}
    p := head
    add := 0
    for l1 != nil || l2 != nil || add > 0 {
        if l1 != nil {
            add += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            add += l2.Val
            l2 = l2.Next
        }
        p.Next = &ListNode{Val: add%10}
        p = p.Next
        add /= 10
    }
    return head.Next
}