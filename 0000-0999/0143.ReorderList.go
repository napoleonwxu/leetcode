/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reorderList(head *ListNode)  {
    if head == nil || head.Next == nil {
        return
    }
    mid, tail := head, head.Next
    for tail != nil && tail.Next != nil {
        mid, tail = mid.Next, tail.Next.Next
    }
    head2 := mid.Next
    mid.Next = nil
    head2 = reverse(head2)
    merge(head, head2)
}

func reverse(head *ListNode) *ListNode {
    var newHead *ListNode
    for head != nil {
        nxt := head.Next
        head.Next = newHead
        newHead = head
        head = nxt
    }
    return newHead
}

func merge(h1, h2 *ListNode) {
    // len(h1) == len(h2) or len(h1) == len(h2)+1
    for h2 != nil {
        n1, n2 := h1.Next, h2.Next
        h1.Next = h2
        h2.Next = n1
        h1, h2 = n1, n2
    }
}