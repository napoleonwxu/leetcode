/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func doubleIt(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    if head.Val >= 5 {
        head = &ListNode{
            Val:  0,
            Next: head,
        }
    }
    for p := head; p != nil; p = p.Next {
        p.Val = 2 * p.Val % 10
        if p.Next != nil && p.Next.Val >= 5 {
            p.Val++
        }
    }
    return head
    /* reverse
    head = reverse(head)
    tmp := head
    for ; tmp != nil; tmp = tmp.Next {
        tmp.Val *= 2
    }
    tmp = head
    for ; tmp != nil; tmp = tmp.Next {
        if tmp.Val >= 10 {
            if tmp.Next == nil {
                tmp.Next = &ListNode{}
            }
            tmp.Next.Val += tmp.Val / 10
            tmp.Val %= 10
        }
    }
    return reverse(head)
    */
}

func reverse(head *ListNode) *ListNode {
    ans := &ListNode{}
    ans = ans.Next
    for head != nil {
        tmp := head.Next
        head.Next = ans
        ans = head
        head = tmp
    }
    return ans
}
