/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
    stack1, stack2 := []*TreeNode{}, []*TreeNode{}
    pushLeft(root1, &stack1)
    pushLeft(root2, &stack2)
    ans := []int{}
    for len(stack1) > 0 || len(stack2) > 0 {
        n1, n2 := len(stack1), len(stack2)
        var node *TreeNode
        if n2 == 0 || (n1 > 0 && stack1[n1-1].Val <= stack2[n2-1].Val) {
            node = stack1[n1-1]
            stack1 = stack1[:n1-1]
            pushLeft(node.Right, &stack1)
        } else {
            node = stack2[n2-1]
            stack2 = stack2[:n2-1]
            pushLeft(node.Right, &stack2)
        }
        ans = append(ans, node.Val)
    }
    return ans
}

func pushLeft(node *TreeNode, stack *[]*TreeNode) {
    for node != nil {
        *stack = append(*stack, node)
        node = node.Left
    }
}