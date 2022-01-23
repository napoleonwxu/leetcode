/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func tree2str(t *TreeNode) string {
    if t == nil {
        return ""
    }
    /* recurive
    s := strconv.Itoa(t.Val)
    if t.Left == nil && t.Right == nil {
        return s
    }
    s += "(" + tree2str(t.Left) + ")"
    if t.Right != nil {
        s += "(" + tree2str(t.Right) + ")"
    }
    return s
    */
    // iterative
    stack := []*TreeNode{t}
    Map := make(map[*TreeNode]int)
    s := ""
    for len(stack) > 0 {
        t := stack[len(stack)-1]
        if Map[t] == 0 {
            Map[t]++
            s += "(" + strconv.Itoa(t.Val)
            if t.Left == nil && t.Right != nil {
                s += "()"
            }
            if t.Right != nil {
                stack = append(stack, t.Right)
            }
            if t.Left != nil {
                stack = append(stack, t.Left)
            }
        } else {
            stack = stack[:len(stack)-1]
            s += ")"
        }
    }
    return s[1:len(s)-1]
}