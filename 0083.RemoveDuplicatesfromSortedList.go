/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func deleteDuplicates(head *ListNode) *ListNode {
    p := head
    for p != nil {
        tmp := p.Next
        for tmp != nil && tmp.Val == p.Val {
            tmp = tmp.Next
        }
        p.Next = tmp
        p = tmp
    }
    return head
}