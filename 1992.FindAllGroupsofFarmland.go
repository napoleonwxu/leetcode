func findFarmland(land [][]int) [][]int {
	ans := [][]int{}
	for i, r := range land {
		for j, v := range r {
			if v == 1 {
				group := []int{i, j, i, j}
				dfs(land, group, i, j)
				ans = append(ans, group)
			}
		}
	}
	return ans
}

func dfs(land [][]int, group []int, i, j int) {
	if i < 0 || i >= len(land) || j < 0 || j >= len(land[0]) || land[i][j] != 1 {
		return
	}
	if i <= group[0] && j <= group[1] {
		group[0], group[1] = i, j
	}
	if i >= group[2] && j >= group[3] {
		group[2], group[3] = i, j
	}
	land[i][j] = -1
	dfs(land, group, i-1, j)
	dfs(land, group, i+1, j)
	dfs(land, group, i, j-1)
	dfs(land, group, i, j+1)
}