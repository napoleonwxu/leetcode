/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func diameterOfBinaryTree(root *TreeNode) int {
    diameter := 0
    dfs(root, &diameter)
    return diameter
}

func dfs(node *TreeNode, diameter *int) int {
    if node == nil {
        return 0
    }
    depthL := dfs(node.Left, diameter)
    depthR := dfs(node.Right, diameter)
    *diameter = max(*diameter, depthL + depthR)
    return max(depthL, depthR) + 1
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}