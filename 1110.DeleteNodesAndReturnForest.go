/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
    Map := make(map[int]bool, len(to_delete))
    for _, d := range to_delete {
        Map[d] = true
    }
    ans := []*TreeNode{}
    dfs(root, true, Map, &ans)
    return ans
}

func dfs(node *TreeNode, isRoot bool, Map map[int]bool, ans *[]*TreeNode) *TreeNode {
    if node == nil {
        return nil
    }
    isDel := Map[node.Val]
    if isRoot && !isDel {
        *ans = append(*ans, node)
    }
    node.Left = dfs(node.Left, isDel, Map, ans)
    node.Right = dfs(node.Right, isDel, Map, ans)
    if isDel {
        return nil
    }
    return node
}