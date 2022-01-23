import "strings"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getDirections(root *TreeNode, startValue int, destValue int) string {
	if root == nil {
		return ""
	}
	startPath, destPath := []byte{}, []byte{}
	findPath(root, startValue, &startPath)
	findPath(root, destValue, &destPath)
	i, j := len(startPath)-1, len(destPath)-1
	for ; i >= 0 && j >= 0 && startPath[i] == destPath[j]; i, j = i-1, j-1 {
	}
	for l, r := 0, j; l < r; l, r = l+1, r-1 {
		destPath[l], destPath[r] = destPath[r], destPath[l]
	}
	return strings.Repeat("U", i+1) + string(destPath[:j+1])
}

func findPath(node *TreeNode, value int, path *[]byte) bool {
	if node.Val == value {
		return true
	}
	if node.Left != nil && findPath(node.Left, value, path) {
		*path = append(*path, 'L')
		return true
	}
	if node.Right != nil && findPath(node.Right, value, path) {
		*path = append(*path, 'R')
		return true
	}
	return false
} 