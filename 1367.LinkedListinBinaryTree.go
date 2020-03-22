/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isSubPath(head *ListNode, root *TreeNode) bool {
    // O(N * min(L,H))
    if root == nil {
        return false
    }
    return dfs(head, root) || isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

func dfs(list *ListNode, node *TreeNode) bool {
    if list == nil {
        return true
    }
    if node == nil || node.Val != list.Val {
        return false
    }
    return dfs(list.Next, node.Left) || dfs(list.Next, node.Right)
}