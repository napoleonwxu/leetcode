/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func postorderTraversal(root *TreeNode) []int {
    ans := []int{}
    stack := []*TreeNode{}
    for len(stack) > 0 || root != nil {
        if root != nil {
            ans = append(ans, root.Val)
            stack = append(stack, root)
            root = root.Right
        } else {
            root = stack[len(stack)-1].Left
            stack = stack[:len(stack)-1]
        }
    }
    for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
        ans[i], ans[j] = ans[j], ans[i]
    }
    return ans
}
