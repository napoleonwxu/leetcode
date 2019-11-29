/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func lcaDeepestLeaves(root *TreeNode) *TreeNode {
    ancestor, _ := dfs(root)
    return ancestor
}

func dfs(node *TreeNode) (*TreeNode, int) {
    if node == nil {
        return nil, 0
    }
    ancestorL, depthL := dfs(node.Left)
    ancestorR, depthR := dfs(node.Right)
    if depthL > depthR {
        return ancestorL, depthL+1
    } else if depthR > depthL {
        return ancestorR, depthR+1
    }
    return node, depthL+1
}