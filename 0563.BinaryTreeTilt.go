/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var ans int

func findTilt(root *TreeNode) int {
	ans = 0
	dfs(root)
	return ans
}

func dfs(node *TreeNode) int {
	if node == nil {
		return 0
	}
	sumL := dfs(node.Left)
	sumR := dfs(node.Right)
	ans += abs(sumL, sumR)
	return node.Val + sumL + sumR
}

func abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}