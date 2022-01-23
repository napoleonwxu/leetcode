/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 type FindElements struct {
    Map map[int]bool
}


func Constructor(root *TreeNode) FindElements {
    Map := make(map[int]bool)
    dfs(root, 0, Map)
    return FindElements{Map: Map}
}


func (this *FindElements) Find(target int) bool {
    return this.Map[target]
}


func dfs(node *TreeNode, val int, Map map[int]bool) {
    if node == nil {
        return
    }
    node.Val = val
    Map[val] = true
    dfs(node.Left, 2*val+1, Map)
    dfs(node.Right, 2*val+2, Map)
}

/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */