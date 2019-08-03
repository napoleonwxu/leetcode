/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	// O(N) + O(H)
    ancestor, _ := dfs(root)
    return ancestor
}

func dfs(node *TreeNode) (ancestor *TreeNode, depth int) {
    ancestor, depth = &TreeNode{}, 0
    if node == nil {
        return
    }
    ancestorL, depthL := dfs(node.Left)
    ancestorR, depthR := dfs(node.Right)
    if depthL > depthR {
        ancestor, depth = ancestorL, depthL+1
    } else if depthL < depthR {
        ancestor, depth = ancestorR, depthR+1
    } else {
        ancestor, depth = node, depthL+1
    }
    return
}