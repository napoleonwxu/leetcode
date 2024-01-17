/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return root
    }
    root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
    /* queue
    queue := []*TreeNode{root}
    for i := 0 ; i < len(queue); i++ {
        node := queue[i]
        if node != nil {
            node.Left, node.Right = node.Right, node.Left
            queue = append(queue, node.Left, node.Right)
        }
    }*/
    return root
}