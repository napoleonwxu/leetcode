/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func longestUnivaluePath(root *TreeNode) int {
    ans := 0
    dfs(root, &ans)
    return ans
}

func dfs(node *TreeNode, ans *int) int {
    if node == nil {
        return 0
    }
    l := dfs(node.Left, ans)
    r := dfs(node.Right, ans)
    lenL, lenR := 0, 0
    if node.Left != nil && node.Left.Val == node.Val {
        lenL = l
    }
    if node.Right != nil && node.Right.Val == node.Val {
        lenR = r
    }
    *ans = max(*ans, lenL + lenR)
    return max(lenL, lenR) + 1
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}