/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func increasingBST(root *TreeNode) *TreeNode {
    head := &TreeNode{}
    cur := head
    var inorder func(*TreeNode)
    inorder = func(node *TreeNode) {
        if node == nil {
            return
        }
        inorder(node.Left)
        node.Left = nil
        cur.Right = node
        cur = node
        inorder(node.Right)
    }
    inorder(root)
    return head.Right
}