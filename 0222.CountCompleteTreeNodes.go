/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func countNodes(root *TreeNode) int {
    if root ==  nil {
        return 0
    }
    // O(logN*logN): logN + (logN)-1 + (logN)-2 + ... + 1 = (1+logN)*logN/2
    hl, hr := 0, 0
    for p := root; p != nil; p = p.Left {
        hl++
    }
    for p := root; p != nil; p = p.Right {
        hr++
    }
    if hl == hr {
        return 2<<uint(hl-1) - 1
    }
    return 1 + countNodes(root.Left) + countNodes(root.Right)
    // O(N)
    //return 1 + countNodes(root.Left) + countNodes(root.Right)
}