/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 func search(root *TreeNode, k int, Map map[int]int) bool {
    if root == nil {
        return false
    }
    _, ok := Map[k-root.Val]
    if ok {
        return true
    }
    Map[root.Val] = 1
    return search(root.Left, k, Map) || search(root.Right, k, Map)
}

func inorder(root *TreeNode, list *[]int) {
    if root == nil {
        return 
    }
    inorder(root.Left, list)
    *list = append(*list, root.Val)
    inorder(root.Right, list)
}

func findTarget(root *TreeNode, k int) bool {
    /* O(n) + O(n)
    Map := make(map[int]int)
    return search(root, k, Map)
    */
    // O(n) + O(n)
    var list []int
    inorder(root, &list)
    left, right := 0, len(list)-1
    for ; left < right; {
        sum := list[left] + list[right]
        if sum == k {
            return true
        } else if sum < k {
            left++
        } else {
            right--
        }
    }
    return false
}