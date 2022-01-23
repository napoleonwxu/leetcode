/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func inorderTraversal(root *TreeNode) []int {
    ans := []int{}
    stack := []*TreeNode{}
    for root != nil || len(stack) > 0 {
        if root != nil {
            stack = append(stack, root)
            root = root.Left
        } else {
            root = stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            ans = append(ans, root.Val)
            root = root.Right
        }
    }
    return ans
}