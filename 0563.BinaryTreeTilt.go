/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func findTilt(root *TreeNode) int {
    tilt := 0
    postorder(root, &tilt)
    return tilt
}

func postorder(node *TreeNode, tilt *int) int {
    if node == nil {
        return 0
    }
    sumL := postorder(node.Left, tilt)
    sumR := postorder(node.Right, tilt)
    *tilt += abs(sumL, sumR)
    return node.Val + sumL + sumR
}

func abs(a, b int) int {
    if a > b {
        return a-b
    }
    return b-a
}