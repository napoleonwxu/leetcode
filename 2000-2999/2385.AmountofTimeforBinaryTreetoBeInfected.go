/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func amountOfTime(root *TreeNode, start int) int {
	m := getMaxVal(root)
	graph := make([][]int, m+1)
	for i := 1; i <= m; i++ {
		graph[i] = make([]int, 0, 3)
	}
	genGraph(root, graph)
	visited := make(map[int]bool, m)
	return dfs(graph, visited, start) - 1
}

func getMaxVal(node *TreeNode) int {
	if node == nil {
		return 0
	}
	m := max(node.Val, getMaxVal(node.Left))
	return max(m, getMaxVal(node.Right))
}

func genGraph(node *TreeNode, graph [][]int) {
	if node.Left != nil {
		graph[node.Val] = append(graph[node.Val], node.Left.Val)
		graph[node.Left.Val] = append(graph[node.Left.Val], node.Val)
		genGraph(node.Left, graph)
	}
	if node.Right != nil {
		graph[node.Val] = append(graph[node.Val], node.Right.Val)
		graph[node.Right.Val] = append(graph[node.Right.Val], node.Val)
		genGraph(node.Right, graph)
	}
}

func dfs(graph [][]int, visited map[int]bool, start int) int {
	if visited[start] {
		return 0
	}
	depth := 0
	visited[start] = true
	for _, nxt := range graph[start] {
		depth = max(depth, dfs(graph, visited, nxt))
	}
	return depth + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}