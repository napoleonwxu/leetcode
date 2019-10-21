/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 type Node struct {
    val, x, y int
}

func verticalTraversal(root *TreeNode) [][]int {
    nodes := []Node{}
    dfs(root, 0, 0, &nodes)
    sort.Slice(nodes, func(i, j int) bool {
        if nodes[i].x == nodes[j].x {
            if nodes[i].y == nodes[j].y {
                return nodes[i].val < nodes[j].val
            }
            return nodes[i].y > nodes[j].y
        }
        return nodes[i].x < nodes[j].x
    })
    ans := [][]int{}
    x, vals := nodes[0].x, []int{}
    for _, node := range nodes {
        if node.x == x {
            vals = append(vals, node.val)
        } else {
            ans = append(ans, vals)
            x, vals = node.x, []int{node.val}
        }
    }
    ans = append(ans, vals)
    return ans
}

func dfs(node *TreeNode, x, y int, nodes *[]Node) {
    if node == nil {
        return
    }
    *nodes = append(*nodes, Node{val: node.Val, x: x, y: y})
    dfs(node.Left, x-1, y-1, nodes)
    dfs(node.Right, x+1, y-1, nodes)
}
