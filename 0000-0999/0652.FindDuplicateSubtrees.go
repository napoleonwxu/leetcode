/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
    Map := make(map[string]int)
    ans := []*TreeNode{}
    // preorder and postorder works here, but inorder does not.
    preorder(root, Map, &ans)
    return ans
}

func preorder(node *TreeNode, Map map[string]int, ans *[]*TreeNode) string {
    if node == nil {
        return "#"
    }
    // wrong answer with string()
    //s := string(node.Val) + "," + dfs(node.Left, Map, ans) + "," + dfs(node.Right, Map, ans)
    s := strconv.Itoa(node.Val) + "," + preorder(node.Left, Map, ans) + "," + preorder(node.Right, Map, ans)
    Map[s]++
    if Map[s] == 2 {
        *ans = append(*ans, node)
    }
    return s
}