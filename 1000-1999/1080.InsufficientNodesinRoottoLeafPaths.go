/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func sufficientSubset(root *TreeNode, limit int) *TreeNode {
    //if root.Left == nil && root.Right == nil {
    if root.Left == root.Right {
        if root.Val >= limit {
            return root
        }
        return nil
    }
    if root.Left != nil {
        root.Left = sufficientSubset(root.Left, limit-root.Val)
    }
    if root.Right != nil {
        root.Right = sufficientSubset(root.Right, limit-root.Val)
    }
    //if root.Left == nil && root.Right == nil {
    if root.Left == root.Right {
        return nil
    }
    return root
}