/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func maxLevelSum(root *TreeNode) int {
    queue := []*TreeNode{root}
    ans, max := 0, 0
    level := 0
    for len(queue) > 0 {
        sum := 0
        nxt := []*TreeNode{}
        for _, node := range queue {
            if node != nil {
                sum += node.Val
                nxt = append(nxt, node.Left, node.Right)
            }
        }
        queue = nxt
        level++
        if sum > max {
            ans, max = level, sum
        }
    }
    return ans
}