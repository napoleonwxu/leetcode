/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func removeZeroSumSublists(head *ListNode) *ListNode {
    Map := make(map[int]*ListNode)
    dummy := &ListNode{}
    dummy.Next = head
    Map[0] = dummy
    sum := 0
    for head != nil {
        sum += head.Val
        if pre, ok := Map[sum]; ok {
            head := pre.Next
            invalid := sum + head.Val
            for invalid != sum {
                delete(Map, invalid)
                head = head.Next
                invalid += head.Val
            }
            pre.Next = head.Next
        } else {
            Map[sum] = head
        }
        head = head.Next
    }
    return dummy.Next
}