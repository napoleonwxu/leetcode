/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 const mod = 1e9 + 7
 var sum, prod int
 
 func maxProduct(root *TreeNode) int {
	 sum = dfs(root, false)
	 prod = 0
	 dfs(root, true)
	 return prod%mod
 }
 
 func dfs(node *TreeNode, calProd bool) int {
	 if node == nil {
		 return 0
	 }
	 tmp := node.Val + dfs(node.Left, calProd) + dfs(node.Right, calProd)
	 if calProd {
		 prod = max(prod, tmp*(sum-tmp))
	 }
	 return tmp
 }
 
 func max(x, y int) int {
	 if x > y {
		 return x
	 }
	 return y
 }