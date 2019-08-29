/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func maxPathSum(root *TreeNode) int {
    ans := math.MinInt32
    dfs(root, &ans)
    return ans
}

func dfs(node *TreeNode, ans *int) int {
    if node == nil {
        return 0
    }
    maxL := dfs(node.Left, ans)
    maxR := dfs(node.Right, ans)
    *ans = max(*ans, node.Val + max(0, maxL) + max(0, maxR))
    return node.Val + max(0, max(maxL, maxR))
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}