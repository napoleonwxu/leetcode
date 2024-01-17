/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    dummy := &ListNode{Next: head}
    pre := dummy
    for i := 1; i < left; i++ {
        pre = pre.Next
    }
    l, r := pre.Next, pre.Next
    for i := left; i < right; i++ {
        r = r.Next
    }
    tail := r.Next
    r.Next = nil
    for l != nil {
        tmp := l.Next
        l.Next = tail
        tail = l
        l = tmp
    }
    pre.Next = r
    return dummy.Next
}