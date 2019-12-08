/*
L1: the distance between the head point and the entry point
L2: the distance between the entry point and the meeting point
C: the length of the cycle
n: the travel times of the fast pointer around the cycle
2*(L1 + L2) = L1 + n*C + L2
L1 + L2 = n*C
L1 = n*C - L2
L1 = (n-1)*C + (C-L2)
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func detectCycle(head *ListNode) *ListNode {
    // O(n) 
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