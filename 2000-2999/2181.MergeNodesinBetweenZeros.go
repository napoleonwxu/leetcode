/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func mergeNodes(head *ListNode) *ListNode {
    p, q := head, head.Next
    for q != nil && q.Next != nil {
        if q.Val == 0 {
            p = p.Next
            p.Val = 0
        }
        p.Val += q.Val
        q = q.Next
    }
    p.Next = nil
    return head
}