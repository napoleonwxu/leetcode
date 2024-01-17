import "sort"

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
*/
func minimumOperations(root *TreeNode) int {
    ans := 0
    nodes := []*TreeNode{root}
    for len(nodes) > 0 {
        nums := make([]int, 0, len(nodes))
        nxt := make([]*TreeNode, 0, 2*len(nodes))
        for _, node := range nodes {
            if node != nil {
                nums = append(nums, node.Val)
                nxt = append(nxt, node.Left, node.Right)
            }
        }
        ans += minSwapSteps(nums)
        nodes = nxt
    }
    return ans
}

func minSwapSteps(nums []int) int {
    n, step := len(nums), 0
    sorted := make([]int, n)
    copy(sorted, nums)
    sort.Ints(sorted)
    idxs := make(map[int]int, n)
    for i, num := range sorted {
        idxs[num] = i
    }
    visited := make([]bool, n)
    for i := range nums {
        cnt := 0
        for !visited[i] && i != idxs[nums[i]] {
            visited[i], i = true, idxs[nums[i]]
            cnt++
        }
        if cnt > 1 {
            step += cnt - 1
        }
    }
    return step
}