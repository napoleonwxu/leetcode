/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    mid, end := head, head
    var prev *ListNode
    for end != nil && end.Next != nil {
        prev = mid
        mid, end = mid.Next, end.Next.Next
    }
    prev.Next = nil
    return merge(sortList(head), sortList(mid))
}

func merge(p1, p2 *ListNode) *ListNode {
    dummy := &ListNode{}
    p := dummy
    for p1 != nil && p2 != nil {
        if p1.Val <= p2.Val {
            p.Next = p1
            p1 = p1.Next
        } else {
            p.Next = p2
            p2 = p2.Next
        }
        p = p.Next
    }
    if p1 != nil {
        p.Next = p1
    } else {
        p.Next = p2
    }
    return dummy.Next
}