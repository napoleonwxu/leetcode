/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// https://leetcode.com/problems/binary-tree-coloring-game/discuss/350570/JavaC%2B%2BPython-Simple-recursion-and-Follow-Up
func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
    var left, right, parent int
    count(root, x, &left, &right)
    parent = n - left - right - 1
    return max(max(left, right), parent) > n/2
}

func count(node *TreeNode, x int, left, right *int) int {
    if node == nil {
        return 0
    }
    cntL := count(node.Left, x, left, right)
    cntR := count(node.Right, x, left, right)
    if node.Val == x {
        *left, *right = cntL, cntR
    }
    return cntL + cntR + 1
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}