/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func rob(root *TreeNode) int {
    robRoot, skipRoot := dfs(root)
    return max(robRoot, skipRoot)
    /* slow
    if root == nil {
        return 0
    }
    val := root.Val
    if root.Left != nil {
        val += rob(root.Left.Left) + rob(root.Left.Right)
    }
    if root.Right != nil {
        val += rob(root.Right.Left) + rob(root.Right.Right)
    }
    return max(val, rob(root.Left)+rob(root.Right))
    */
}

func dfs(node *TreeNode) (int, int) {
    if node == nil {
        return 0, 0
    }
    robLeft, skipLeft := dfs(node.Left)
    robRight, skipRight := dfs(node.Right)
    robNode, skipNode := 0, 0
    robNode = node.Val + skipLeft + skipRight
    skipNode = max(robLeft, skipLeft) + max(robRight, skipRight)
    return robNode, skipNode
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}