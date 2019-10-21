/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// O(N*logK), 4-8ms < 99.75%; O(1), 5.3MB < 100% 
func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }
    return divideConquer(lists, 0, len(lists)-1)
    /*
    n := len(lists)
    if n == 0 {
        return nil
    }
    for n > 1 {
        for i := 0; i < n/2; i++ {
            lists[i] = mergeTwoLists(lists[i], lists[n-i-1])
        }
        n = (n+1) / 2
    }
    return lists[0]
    */
}

func divideConquer(lists []*ListNode, left int, right int) *ListNode {
    if left == right {
        return lists[left]
    }
    mid := (left + right)/2
    l1 := divideConquer(lists, left, mid)
    l2 := divideConquer(lists, mid+1, right)
    return mergeTwoLists(l1, l2)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    head := &ListNode{}
    p := head
    for l1 != nil && l2 != nil {
        if l1.Val <= l2.Val {
            p.Next = l1
            l1 = l1.Next
        } else {
            p.Next = l2
            l2 = l2.Next
        }
        p = p.Next
    }
    if l1 != nil {
        p.Next = l1
    } else {
        p.Next = l2
    }
    return head.Next
}

// O(N*K), 108ms < 34.41%; O(1), 5.3MB < 100%
// N = sum(len(lists[i])), K = len(lists)
/*
func mergeKLists(lists []*ListNode) *ListNode {
    n := len(lists)
    if n == 0 {
        return nil
    }
    ans := lists[0]
    for i := 1; i < n; i++ {
        ans = mergeTwoLists(ans, lists[i])
    }
    return ans
}
*/