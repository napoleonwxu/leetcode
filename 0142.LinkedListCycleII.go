/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func detectCycle(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow, fast = slow.Next, fast.Next.Next
        if slow == fast {
            slow2 := head
            for slow != slow2 {
                slow, slow2 = slow.Next, slow2.Next
            }
            return slow
        }
    }
    return nil
}