func checkValidGrid(grid [][]int) bool {
	n := len(grid)
	if n <= 0 || grid[0][0] != 0 {
		return false
	}
	locs := make([][]int, n*n)
	for i, row := range grid {
		for j, step := range row {
			locs[step] = []int{i, j}
		}
	}
	for i := 1; i < len(locs); i++ {
		if abs((locs[i][0]-locs[i-1][0])*(locs[i][1]-locs[i-1][1])) != 2 {
			return false
		}
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}