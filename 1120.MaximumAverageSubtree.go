/*
Given the root of a binary tree, find the maximum average value of any subtree of that tree.
(A subtree of a tree is any node of that tree plus all its descendants. The average value of a tree is the sum of its values, divided by the number of nodes.)
Note:
The number of nodes in the tree is between 1 and 5000.
Each node will have a value between 0 and 100000.
Answers will be accepted as correct if they are within 10^-5 of the correct answer.
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func maximumAverageSubtree(root *TreeNode) float64 {
    maxAvg := 0.0
    dfs(root, &maxAvg)
    return maxAvg
}

func dfs(node *TreeNode, maxAvg *float64) (sum, cnt int) {
    sum, cnt = 0, 0
    if node == nil {
        return
    }
    sumL, cntL := dfs(node.Left, maxAvg)
    sumR, cntR := dfs(node.Right, maxAvg)
    sum, cnt = sumL+sumR+node.Val, cntL+cntR+1
    *maxAvg = max(*maxAvg, float64(sum)/float64(cnt))
    return
}

func max(x, y float64) float64 {
    if x > y {
        return x
    }
    return y
}