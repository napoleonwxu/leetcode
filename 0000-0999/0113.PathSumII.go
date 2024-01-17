/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
    ans := [][]int{}
    dfs(root, sum, []int{}, &ans)
    return ans
}

func dfs(node *TreeNode, sum int, path []int, ans *[][]int) {
    if node == nil {
        return
    }
    if node.Left == nil && node.Right == nil && node.Val == sum {
        path = append(path, node.Val)
        tmp := make([]int, len(path))
        copy(tmp, path)
        *ans = append(*ans, tmp)
        return
    }
    dfs(node.Left, sum-node.Val, append(path, node.Val), ans)
    dfs(node.Right, sum-node.Val, append(path, node.Val), ans)
}