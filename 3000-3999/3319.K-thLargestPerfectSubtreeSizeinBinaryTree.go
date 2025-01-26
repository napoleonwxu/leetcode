/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func kthLargestPerfectSubtree(root *TreeNode, k int) int {
    size := make([]int, 0)
    dfs(root, &size)
    sort.Ints(size)
    if len(size) >= k {
        return size[len(size)-k]
    }
    return -1
}

func dfs(node *TreeNode, size *[]int) (bool, int) {
    if node == nil {
        return true, 0
    }
    l, lc := dfs(node.Left, size)
    r, rc := dfs(node.Right, size)
    perfect := l && r && lc == rc
    if perfect {
        *size = append(*size, lc+rc+1)
    }
    return perfect, lc + rc + 1
}