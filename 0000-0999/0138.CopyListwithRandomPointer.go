/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    if head == nil {
        return nil
    }
    p := head
    for p != nil {
        q := &Node{Val: p.Val, Next: p.Next}
        p.Next = q
        p = q.Next
    }
    p = head
    for p != nil {
        q := p.Next
        if p.Random != nil {
            q.Random = p.Random.Next
        }
        p = q.Next
    }
    newHead := head.Next
    p, q := head, newHead
    for q.Next != nil {
        p.Next = q.Next
        p = p.Next
        q.Next = p.Next
        q = q.Next
    }
    p.Next = nil
    return newHead
}