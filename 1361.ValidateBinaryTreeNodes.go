/* This solution is wrong but accepted
3
[1,0,-1]
[-1,-1,-1]
*/
func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
    nodes := make(map[int]bool)
    child := append(leftChild, rightChild...)
    for _, n := range child {
        if n == -1 {
            continue
        }
        if nodes[n] {
            return false
        }
        nodes[n] = true
    }
    if len(nodes) != n-1 {
        return false
    }
    return true
}