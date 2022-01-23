/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil {
        return head
    }
    tail, n := head, 1
    for tail.Next != nil {
        tail = tail.Next
        n++
    }
    k %= n
    if k == 0 {
        return head
    }
    tail.Next = head
    for i := 0; i < n-k; i++ {
        tail = tail.Next
    }
    newHead := tail.Next
    tail.Next = nil
    return newHead
}