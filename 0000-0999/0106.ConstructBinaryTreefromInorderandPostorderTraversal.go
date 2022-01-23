/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func buildTree(inorder []int, postorder []int) *TreeNode {
    n := len(postorder)
    if n == 0 {
        return nil
    }
    i := index(inorder, postorder[n-1])
    node := &TreeNode{Val: postorder[n-1]}
    node.Left = buildTree(inorder[:i], postorder[:i])
    node.Right = buildTree(inorder[i+1:], postorder[i:n-1])
    return node
}

func index(nums []int, val int) int {
    for i, num := range nums {
        if num == val {
            return i
        }
    }
    return -1
}