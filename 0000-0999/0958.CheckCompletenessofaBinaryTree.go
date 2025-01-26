/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isCompleteTree(root *TreeNode) bool {
    // BFS
    queue := []*TreeNode{root}
    i := 0
    for ; queue[i] != nil; i++ {
        queue = append(queue, queue[i].Left, queue[i].Right)
    }
    for ; i < len(queue); i++ {
        if queue[i] != nil {
            return false
        }
    }
    return true
}