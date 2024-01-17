/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
*/
func reverseOddLevels(root *TreeNode) *TreeNode {
    dfs(root.Left, root.Right, 1)
    return root
}

func dfs(l, r *TreeNode, depth int) {
    if l == nil || r == nil {
        return
    }
    if depth%2 == 1 {
        l.Val, r.Val = r.Val, l.Val
    }
    dfs(l.Left, r.Right, depth+1)
    dfs(l.Right, r.Left, depth+1)
}