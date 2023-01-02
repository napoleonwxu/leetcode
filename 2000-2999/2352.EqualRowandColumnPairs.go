import (
	"strconv"
	"strings"
)

func equalPairs(grid [][]int) int {
	n := len(grid)
	m := make(map[string]int, n)
	for i := 0; i < n; i++ {
		strs := make([]string, n)
		for j := 0; j < n; j++ {
			strs[j] = strconv.Itoa(grid[i][j])
		}
		m[strings.Join(strs, "-")]++
	}
	cnt := 0
	for j := 0; j < n; j++ {
		strs := make([]string, n)
		for i := 0; i < n; i++ {
			strs[i] = strconv.Itoa(grid[i][j])
			cnt += m[strings.Join(strs, "-")]
		}
	}
	return cnt
}