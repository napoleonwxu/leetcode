/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func smallestFromLeaf(root *TreeNode) string {
    ans := string('z'+1)
    dfs(root, "", &ans)
    return ans
}

func dfs(node *TreeNode, s string, ans *string) {
    if node == nil {
        return
    }
    s = string('a' + node.Val) + s
    if node.Left == nil && node.Right == nil && s < *ans {
        *ans = s
    }
    dfs(node.Left, s, ans)
    dfs(node.Right, s, ans)
}