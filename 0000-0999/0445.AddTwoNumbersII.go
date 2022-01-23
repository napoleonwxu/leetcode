/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    s1, s2 := toNum(l1), toNum(l2)
    c := 0
    var head *ListNode
    for len(s1) > 0 || len(s2) > 0 || c > 0 {
        if len(s1) > 0 {
            c += s1[len(s1)-1]
            s1 = s1[:len(s1)-1]
        }
        if len(s2) > 0 {
            c += s2[len(s2)-1]
            s2 = s2[:len(s2)-1]
        }
        node := &ListNode{Val: c%10, Next: head}
        head = node
        c /= 10
    }
    return head
}

func toNum(l *ListNode) []int {
    stack := []int{}
    for l != nil {
        stack = append(stack, l.Val) 
        l = l.Next
    }
    return stack
}