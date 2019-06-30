/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    leaf1 := [100]int{}
    leaf2 := [100]int{}
    p1 := leaf1[:0]
    p2 := leaf2[:0]
    dfs(root1, &p1)
    dfs(root2, &p2)
    return leaf1 == leaf2
}

func dfs(root *TreeNode, p *[]int) {
    if root == nil {
        return
    }
    if root.Left == nil && root.Right == nil {
        *p = append(*p, root.Val)
        return
    }
    dfs(root.Left, p)
    dfs(root.Right, p)
}