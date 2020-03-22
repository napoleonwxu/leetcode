/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 var sum int

 func sumEvenGrandparent(root *TreeNode) int {
	 sum = 0
	 dfs(root, 1, 1)
	 return sum
 }
 
 func dfs(node *TreeNode, pVal, gpVal int) {
	 if node == nil {
		 return
	 }
	 if gpVal&1 == 0 {
		 sum += node.Val
	 }
	 dfs(node.Left, node.Val, pVal)
	 dfs(node.Right, node.Val, pVal)
 }