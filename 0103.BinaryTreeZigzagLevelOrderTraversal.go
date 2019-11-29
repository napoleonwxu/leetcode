/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 var ans [][]int

 func zigzagLevelOrder(root *TreeNode) [][]int {
     ans = [][]int{}
     dfs(root, 0)
     for i := 1; i < len(ans); i += 2 {
         reverse(ans[i])
     }
     return ans
 }
 
 func dfs(node *TreeNode, i int) {
     if node == nil {
         return
     }
     if len(ans) == i {
         ans = append(ans, []int{node.Val})
     } else {
         ans[i] = append(ans[i], node.Val)
     }
     dfs(node.Left, i+1)
     dfs(node.Right, i+1)
 }
 
 func reverse(nums []int) {
     for l, r := 0, len(nums)-1; l < r; l, r = l+1, r-1 {
         nums[l], nums[r] = nums[r], nums[l]
     }
 }