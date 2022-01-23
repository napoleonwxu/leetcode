/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	// Bottom-up, O(n)
    return dfs(root) != -1
}

func dfs(node *TreeNode) int {
    if node == nil {
        return 0
    }
    depthL, depthR := dfs(node.Left), dfs(node.Right)
    if depthL == -1 || depthR == -1 || abs(depthL-depthR) > 1 {
        return -1
    }
    return max(depthL, depthR) + 1
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}