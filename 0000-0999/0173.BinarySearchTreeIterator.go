/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 type BSTIterator struct {
    stack []*TreeNode
}


func Constructor(root *TreeNode) BSTIterator {
    bst := BSTIterator{stack: []*TreeNode{}}
    bst.pushLeft(root)
    return bst
}


/** @return the next smallest number */
func (this *BSTIterator) Next() int {
    n := len(this.stack)
    node := this.stack[n-1]
    this.stack = this.stack[:n-1]
    this.pushLeft(node.Right)
    return node.Val
}


/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
    return len(this.stack) > 0
}

func (this *BSTIterator) pushLeft(node *TreeNode) {
    for node != nil {
        this.stack = append(this.stack, node)
        node = node.Left
    }
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */