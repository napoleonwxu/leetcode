/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func largestValues(root *TreeNode) []int {
    ans := []int{}
    // BFS
    if root == nil {
        return ans
    }
    level := []*TreeNode{root}
    for len(level) > 0 {
        next := []*TreeNode{}
        max := math.MinInt32
        for _, node := range level {
            if node.Val > max {
                max = node.Val
            }
            if node.Left != nil {
                next = append(next, node.Left)
            }
            if node.Right != nil {
                next = append(next, node.Right)
            }
        }
        ans = append(ans, max)
        level = next
    }
    // DFS, cool
    //dfs(root, &ans, 0)
    return ans
}

func dfs(node *TreeNode, ans *[]int, depth int) {
    if node == nil {
        return
    }
    if depth == len(*ans) {
        *ans = append(*ans, node.Val)
    } else {
        if node.Val > (*ans)[depth] {
            (*ans)[depth] = node.Val
        }
    }
    dfs(node.Left, ans, depth+1)
    dfs(node.Right, ans, depth+1)
}
