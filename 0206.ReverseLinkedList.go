/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reverseList(head *ListNode) *ListNode {
    var p *ListNode
    //p = nil
    for head != nil {
        tmp := head.Next
        head.Next = p
        p = head
        head = tmp
    }
    return p
}