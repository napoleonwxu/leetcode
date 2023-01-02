import "sort"

func deleteGreatestValue(grid [][]int) int {
	for _, g := range grid {
		sort.Ints(g)
	}
	ans := 0
	for j := len(grid[0]) - 1; j >= 0; j-- {
		tmp := 0
		for i := 0; i < len(grid); i++ {
			tmp = max(tmp, grid[i][j])
		}
		ans += tmp
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}