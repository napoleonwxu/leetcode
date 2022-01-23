/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
	return dfs(root, root.Val)
}

func dfs(node *TreeNode, maxVal int) int {
	if node == nil {
		return 0
	}
	cnt := 0
	if node.Val >= maxVal {
		cnt++
		maxVal = node.Val
	}
	return cnt + dfs(node.Left, maxVal) + dfs(node.Right, maxVal)
}