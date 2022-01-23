/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isSubtree(s *TreeNode, t *TreeNode) bool {
    if s == nil {
        return false
    }
    if isSame(s, t) {
        return true
    }
    return isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func isSame(s, t *TreeNode) bool {
    if s == nil && t == nil {
        return true
    }
    if s == nil || t == nil || s.Val != t.Val {
        return false
    }
    return isSame(s.Left, t.Left) && isSame(s.Right, t.Right)
}