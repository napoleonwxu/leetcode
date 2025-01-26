/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func deleteDuplicates(head *ListNode) *ListNode {
    dummy := &ListNode{}
    dummy.Next = head
    pre, cur := dummy, head
    for cur != nil {
        for cur.Next != nil && cur.Next.Val == cur.Val {
            cur = cur.Next
        }
        if pre.Next == cur {
            pre = pre.Next
        } else {
            pre.Next = cur.Next
        }
        cur = cur.Next
    }
    return dummy.Next
}