/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func deleteNode(root *TreeNode, key int) *TreeNode {
    if root == nil {
        return root
    }
    if key < root.Val {
        root.Left = deleteNode(root.Left, key)
    } else if key > root.Val {
        root.Right = deleteNode(root.Right, key)
    } else {
        if root.Left == nil {
            return root.Right
        }
        if root.Right == nil {
            return root.Left
        }
        minRight := findMin(root.Right)
        root.Val = minRight
        root.Right = deleteNode(root.Right, minRight)
    }
    return root
}

func findMin(node *TreeNode) int {
    for node.Left != nil {
        node = node.Left
    }
    return node.Val
}