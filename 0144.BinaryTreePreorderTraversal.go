/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func preorderTraversal(root *TreeNode) []int {
    stack := []*TreeNode{}
    ans := []int{}
    for root != nil || len(stack) > 0 {
        if root != nil {
            ans = append(ans, root.Val)
            stack = append(stack, root)
            root = root.Left
        } else {
            root = stack[len(stack)-1].Right
            stack = stack[:len(stack)-1]
        }
    }
    return ans
}