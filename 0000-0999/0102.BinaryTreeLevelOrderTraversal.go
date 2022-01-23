/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }
	ans := [][]int{}
	// DFS
    dfs(root, 0, &ans)
    /* BFS
    queue := []*TreeNode{root}
    for len(queue) > 0 {
        nums := make([]int, 0, len(queue))
        nxt := make([]*TreeNode, 0, 2*len(queue))
        for _, node := range queue {
            nums = append(nums, node.Val)
            if node.Left != nil {
                nxt = append(nxt, node.Left)
            }
            if node.Right != nil {
                nxt = append(nxt, node.Right)
            }
        }
        ans = append(ans, nums)
        queue = nxt
    }*/
    return ans
}

func dfs(node *TreeNode, depth int, ans *[][]int) {
    if node == nil {
        return
    }
    if depth == len(*ans) {
        *ans = append(*ans, []int{})
    }
    (*ans)[depth] = append((*ans)[depth], node.Val)
    dfs(node.Left, depth+1, ans)
    dfs(node.Right, depth+1, ans)
}