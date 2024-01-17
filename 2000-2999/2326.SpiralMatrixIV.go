/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
*/
func spiralMatrix(m int, n int, head *ListNode) [][]int {
    ans := make([][]int, m)
    for i := 0; i < m; i++ {
        ans[i] = make([]int, n)
        for j := 0; j < n; j++ {
            ans[i][j] = -1
        }
    }
    upper, bottom := 0, m-1
    left, right := 0, n-1
    for head != nil {
        for j := left; head != nil && j <= right; j++ {
            ans[upper][j] = head.Val
            head = head.Next
        }
        for i := upper + 1; head != nil && i <= bottom; i++ {
            ans[i][right] = head.Val
            head = head.Next
        }
        for j := right - 1; head != nil && j >= left; j-- {
            ans[bottom][j] = head.Val
            head = head.Next
        }
        for i := bottom - 1; head != nil && i > upper; i-- {
            ans[i][left] = head.Val
            head = head.Next
        }
        upper, bottom = upper+1, bottom-1
        left, right = left+1, right-1
    }
    return ans
}