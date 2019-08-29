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
    cur := head
    for i := 0; i < k; i++ {
        if cur == nil {
            return head
        }
        cur = cur.Next
    }
    nxt := reverseKGroup(cur, k)
    for ; k > 1; k-- {
        tmp := head.Next
        head.Next = nxt
        nxt = head
        head = tmp
    }
    head.Next = nxt
    return head
}