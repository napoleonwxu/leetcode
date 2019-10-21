/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func hasCycle(head *ListNode) bool {
    // O(n) + O(1)
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow, fast = slow.Next, fast.Next.Next
        if slow == fast {
            return true
        }
    }
    return false
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