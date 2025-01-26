/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func modifiedList(nums []int, head *ListNode) *ListNode {
    m := make(map[int]bool, len(nums))
    for _, num := range nums {
        m[num] = true
    }
    dummy := ListNode{}
    pre := &dummy
    for ; head != nil; head = head.Next {
        if !m[head.Val] {
            pre.Next = head
            pre = pre.Next
        }
    }
    pre.Next = nil
    return dummy.Next
}