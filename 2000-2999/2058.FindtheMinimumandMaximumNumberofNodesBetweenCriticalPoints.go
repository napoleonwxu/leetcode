/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nodesBetweenCriticalPoints(head *ListNode) []int {
	idxs := []int{}
	pre, idx := head.Val, 1
	for p := head.Next; p.Next != nil; p = p.Next {
		if (p.Val > pre && p.Val > p.Next.Val) || (p.Val < pre && p.Val < p.Next.Val) {
			idxs = append(idxs, idx)
		}
		pre = p.Val
		idx++
	}
	n := len(idxs)
	if n < 2 {
		return []int{-1, -1}
	}
	ans := []int{idxs[n-1] - idxs[0], idxs[n-1] - idxs[0]}
	for i := 1; i < n; i++ {
		ans[0] = min(ans[0], idxs[i]-idxs[i-1])
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}