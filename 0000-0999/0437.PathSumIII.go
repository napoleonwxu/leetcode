/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func pathSum(root *TreeNode, sum int) int {
    // O(n) + O(n)
    Map := make(map[int]int)
    Map[0] = 1
    return dfs(root, sum, 0, Map)
    /* O(n^2) ~ O(nlgn) + O(n) ~ O(lgn)
    if root == nil {
        return 0
    }
    return pathSumFrom(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
    */
}

func dfs(node *TreeNode, sum, curSum int, Map map[int]int) int {
    if node == nil {
        return 0
    }
    curSum += node.Val
    cnt := Map[curSum-sum]
    Map[curSum]++
    cnt += dfs(node.Left, sum, curSum, Map) + dfs(node.Right, sum, curSum, Map)
    Map[curSum]--
    return cnt
}

// O(n) ~ O(lgn)
func pathSumFrom(node *TreeNode, sum int) int {
    if node == nil {
        return 0
    }
    cnt := 0
    if node.Val == sum {
        cnt++
    }
    return cnt + pathSumFrom(node.Left, sum-node.Val) + pathSumFrom(node.Right, sum-node.Val)
}