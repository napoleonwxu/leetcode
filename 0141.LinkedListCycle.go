/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func hasCycle(head *ListNode) bool {
    // O(n) + O(1)
    if head == nil || head.Next == nil {
        return false
    }
    slow, fast := head, head.Next
    for slow != fast {
        if fast == nil || fast.Next == nil {
            return false
        }
        slow, fast = slow.Next, fast.Next.Next
    }
    return true
    /* O(n) + O(n)
    Map := make(map[*ListNode]bool)
    p := head
    for p != nil {
        if Map[p] {
            return true
        }
        Map[p] = true
        p = p.Next
    }
    return false
    */
}