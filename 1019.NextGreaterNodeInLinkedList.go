/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func nextLargerNodes(head *ListNode) []int {
    nums, n := []int{}, 0
    for ; head != nil; head = head.Next {
        nums = append(nums, head.Val)
        n++
    }
    ans := make([]int, n)
    stack, idx := make([]int, n), -1
    for i, num := range nums {
        for ; idx >= 0 && nums[stack[idx]] < num; idx-- {
            ans[stack[idx]] = num
        }
        idx++
        stack[idx] = i
    }
    return ans
}