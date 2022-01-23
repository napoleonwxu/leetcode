/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func getMinimumDifference(root *TreeNode) int {
    // in-order traversal and compare, O(n) + O(h). cool!
    min, pre := 1024, -1024
    in_order(root, &min, &pre)
    return min
    /* in-order traversal and save then compare, O(n+n) + O(n)
    nums := []int{}
    inorder(root, &nums)
    min := nums[1] - nums[0]
    for i := 2; i < len(nums); i++ {
        diff := nums[i] - nums[i-1]
        if diff < min {
            min = diff
        }
    }
    return min
    */
}

func in_order(node *TreeNode, min, pre *int) {
    if node == nil {
        return
    }
    in_order(node.Left, min, pre)
    diff := node.Val - *pre
    if diff < *min {
        *min = diff
    }
    *pre = node.Val
    in_order(node.Right, min, pre)
}

func inorder(node *TreeNode, nums *[]int) {
    if node.Left != nil {
        inorder(node.Left, nums)
    }
    *nums = append(*nums, node.Val)
    if node.Right != nil {
        inorder(node.Right, nums)
    }
}
