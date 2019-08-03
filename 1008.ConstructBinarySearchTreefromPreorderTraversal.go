/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func bstFromPreorder(preorder []int) *TreeNode {
    // O(n), cool !
    i, n := 0, len(preorder)
    var buildBST func(int) *TreeNode
    buildBST = func(bound int) *TreeNode {
        if i == n || preorder[i] > bound {
            return nil
        }
        root := &TreeNode{Val: preorder[i]}
        i++
        root.Left = buildBST(root.Val)
        root.Right = buildBST(bound)
        return root
    }
	return buildBST(1<<32)
	
    /* O(n^2)
    if len(preorder) == 0 {
        return nil
    }
    root := &TreeNode{Val: preorder[0]}
    i := 1
    for ; i < len(preorder) && preorder[0] > preorder[i]; i++ {}
    root.Left = bstFromPreorder(preorder[1:i])
    root.Right = bstFromPreorder(preorder[i:])
    return root
    */
}