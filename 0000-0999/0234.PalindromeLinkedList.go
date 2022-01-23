/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func isPalindrome(head *ListNode) bool {
    slow, fast := head, head
    // List will be modified, O(n) + O(1)
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    tail := reverse(slow)
    for head != nil && tail != nil {
        if head.Val != tail.Val {
            return false
        }
        head, tail = head.Next, tail.Next
    }
    /* List will not be modified, O(n) + O(n)
    stack := []int{}
    for fast != nil && fast.Next != nil {
        stack = append(stack, slow.Val)
        slow = slow.Next
        fast = fast.Next.Next
    }
    idx := len(stack) - 1
    if fast != nil { // odd
        slow = slow.Next
    }
    for slow != nil {
        if slow.Val != stack[idx] {
            return false
        }
        slow = slow.Next
        idx--
    }
    */
    return true
}

func reverse(head *ListNode) *ListNode {
    var newHead *ListNode
    for head != nil {
        tmp := head.Next
        head.Next = newHead
        newHead = head
        head = tmp
    }
    return newHead
}