/*
Given two binary search trees, return True if and only if there is a node in the first tree and a node in the second tree whose values sum up to a given integer target.

Example 1:
Input: root1 = [2,1,4], root2 = [1,0,3], target = 5
Output: true
Explanation: 2 and 3 sum up to 5.

Example 2:
Input: root1 = [0,-10,10], root2 = [5,1,7,0,2], target = 18
Output: false

Constraints:
Each tree has at most 5000 nodes.
-10^9 <= target, node.val <= 10^9
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
 func twoSumBSTs(root1 *TreeNode, root2 *TreeNode, target int) bool {
    if root1 == nil {
        return false
    }
    if search(root2, target-root1.Val) {
        return true
    }
    return twoSumBSTs(root1.Left, root2, target) || twoSumBSTs(root1.Right, root2, target)
}

func search(node *TreeNode, target int) bool {
    if node == nil {
        return false
    }
    if node.Val == target {
        return true
    } else if target < node.Val {
        return search(node.Left, target)
    }
    return search(node.Right, target)
}