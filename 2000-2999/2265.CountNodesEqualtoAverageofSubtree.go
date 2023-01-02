/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 var ans int

 func averageOfSubtree(root *TreeNode) int {
	 ans = 0
	 dfs(root)
	 return ans
 }
 
 func dfs(node *TreeNode) (int, int) {
	 if node == nil {
		 return 0, 0
	 }
	 sumL, cntL := dfs(node.Left)
	 sumR, cntR := dfs(node.Right)
	 sum, cnt := sumL+sumR+node.Val, cntL+cntR+1
	 if sum/cnt == node.Val {
		 ans++
	 }
	 return sum, cnt
 }