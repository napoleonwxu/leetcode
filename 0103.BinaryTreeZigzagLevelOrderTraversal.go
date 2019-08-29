/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func zigzagLevelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }
    ans := [][]int{}
    queue := []*TreeNode{root}
    for len(queue) > 0 {
        vals := []int{}
        nxt := []*TreeNode{}
        for _, node := range queue {
            vals = append(vals, node.Val)
            if node.Left != nil {
                nxt = append(nxt, node.Left)
            }
            if node.Right != nil {
                nxt = append(nxt, node.Right)
            }
        }
        ans = append(ans, vals)
        queue = nxt
    }
    for i := 1; i < len(ans); i += 2 {
        reverse(ans[i])
    }
    return ans
}

func reverse(vals []int) {
    for i, j := 0, len(vals)-1; i < j; i, j= i+1, j-1 {
        vals[i], vals[j] = vals[j], vals[i]
    }
}