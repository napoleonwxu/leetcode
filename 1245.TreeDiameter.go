/*
Given an undirected tree, return its diameter: the number of edges in a longest path in that tree.
The tree is given as an array of edges where edges[i] = [u, v] is a bidirectional edge between nodes u and v.  Each node has labels in the set {0, 1, ..., edges.length}.

Example 1:
Input: edges = [[0,1],[0,2]]
Output: 2
Explanation: 
A longest path of the tree is the path 1 - 0 - 2.

Example 2:
Input: edges = [[0,1],[1,2],[2,3],[1,4],[4,5]]
Output: 4
Explanation: 
A longest path of the tree is the path 3 - 2 - 1 - 4 - 5.

Constraints:
0 <= edges.length < 10^4
edges[i][0] != edges[i][1]
0 <= edges[i][j] <= edges.length
The given edges form an undirected tree.
*/

var graph map[int][]int
var ans int

func treeDiameter(edges [][]int) int {
    graph = make(map[int][]int)
    for _, edge := range edges {
        graph[edge[0]] = append(graph[edge[0]], edge[1])
        graph[edge[1]] = append(graph[edge[1]], edge[0])
    }
    ans = 0
    dfs(0, -1)
    return ans
}

func dfs(cur, pre int) int {
    d1, d2 := 0, 0
    for _, nxt := range graph[cur] {
        if nxt == pre {
            continue
        }
        d := dfs(nxt, cur)
        if d > d1 {
            d1, d2 = d, d1
        } else if d > d2 {
            d2 = d
        }
    }
    ans = max(ans, d1+d2)
    return d1+1
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}