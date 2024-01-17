import "sort"

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
*/
func kthLargestLevelSum(root *TreeNode, k int) int64 {
    if root == nil {
        return -1
    }
    sums := []int64{}
    nodes := []*TreeNode{root}
    for len(nodes) > 0 {
        sum := int64(0)
        nxt := make([]*TreeNode, 0, 2*len(nodes))
        for _, node := range nodes {
            if node != nil {
                sum += int64(node.Val)
                if node.Left != nil {
                    nxt = append(nxt, node.Left)
                }
                if node.Right != nil {
                    nxt = append(nxt, node.Right)
                }
            }
        }
        sums = append(sums, sum)
        nodes = nxt
    }
    if len(sums) < k {
        return -1
    }
    sort.Slice(sums, func(i, j int) bool {
        return sums[i] > sums[j]
    })
    return sums[k-1]
}