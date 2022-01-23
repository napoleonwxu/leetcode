/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func rightSideView(root *TreeNode) []int {
    ans := []int{}
    dfs(root, 0, &ans)
    /* BFS
    if root == nil {
        return nil
    }
    ans := []int{}
    queue := []*TreeNode{root}
    for len(queue) > 0 {
        ans = append(ans, queue[len(queue)-1].Val)
        nxt := []*TreeNode{}
        for _, node := range queue {
            if node.Left != nil {
                nxt = append(nxt, node.Left)
            }
            if node.Right != nil {
                nxt = append(nxt, node.Right)
            }
        }
        queue = nxt
    } */
    return ans
}

func dfs(node *TreeNode, depth int, ans *[]int) {
    if node == nil {
        return
    }
    if len(*ans) == depth {
        *ans = append(*ans, node.Val)
    }
    dfs(node.Right, depth+1, ans)
    dfs(node.Left, depth+1, ans)
}