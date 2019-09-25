/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func pathSum(root *TreeNode, sum int) [][]int {
    if root == nil {
        return nil
    }
    ans := [][]int{}
    dfs(root, sum, []int{}, &ans)
    return ans
}

func dfs(node *TreeNode, sum int, path []int, ans *[][]int) {
    if node.Left == nil && node.Right == nil && sum == node.Val {
        tmp := make([]int, len(path)+1)
        copy(tmp, path)
        tmp[len(path)] = node.Val
        *ans = append(*ans, tmp)
    }
    if node.Left != nil {
        dfs(node.Left, sum-node.Val, append(path, node.Val), ans)
    }
    if node.Right != nil {
        dfs(node.Right, sum-node.Val, append(path, node.Val), ans)
    }
}