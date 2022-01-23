/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 var vals []int

 func balanceBST(root *TreeNode) *TreeNode {
	 vals = []int{}
	 inorder(root)
	 return toBST(0, len(vals)-1)
 }
 
 func inorder(node *TreeNode) {
	 if node == nil {
		 return
	 }
	 inorder(node.Left)
	 vals = append(vals, node.Val)
	 inorder(node.Right)
 }
 
 func toBST(start, end int) *TreeNode {
	 if start > end {
		 return nil
	 }
	 mid := (start+end) >> 1
	 node := &TreeNode{Val: vals[mid]}
	 node.Left = toBST(start, mid-1)
	 node.Right = toBST(mid+1, end)
	 return node
 }