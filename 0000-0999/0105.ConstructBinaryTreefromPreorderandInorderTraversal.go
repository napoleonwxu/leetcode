/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(inorder) == 0 {
        return nil
    }
    idx := index(inorder, preorder[0])
    root := &TreeNode{Val: preorder[0]}
    root.Left = buildTree(preorder[1:idx+1], inorder[:idx])
    root.Right = buildTree(preorder[idx+1:], inorder[idx+1:])
    return root
}

func index(inorder []int, val int) int {
    for i, v := range inorder {
        if v == val {
            return i
        }
    }
    return -1
}