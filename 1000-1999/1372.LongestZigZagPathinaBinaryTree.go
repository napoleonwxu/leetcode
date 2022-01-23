/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 func longestZigZag(root *TreeNode) int {
    return dfs(root)[2]
}

func dfs(node *TreeNode) []int {
    if node == nil {
        return []int{-1, -1, -1}
    }
    left, right := dfs(node.Left), dfs(node.Right)
    ans := max(max(1+left[1], 1+right[0]), max(left[2], right[2]))
    return []int{1+left[1], 1+right[0], ans}
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}