/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func constructFromPrePost(pre []int, post []int) *TreeNode {
    n := len(pre)
    if n == 0 {
        return nil
    }
    root := &TreeNode{Val: pre[0]}
    if n == 1 {
        return root
    }
    i := idx(post, pre[1])
    root.Left = constructFromPrePost(pre[1:i+2], post[:i+1])
    root.Right = constructFromPrePost(pre[i+2:], post[i+1:n-1])
    return root
}

func idx(nums []int, target int) int {
    for i, num := range nums {
        if num == target {
            return i
        }
    }
    return -1
}