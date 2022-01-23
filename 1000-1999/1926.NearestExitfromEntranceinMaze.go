func nearestExit(maze [][]byte, entrance []int) int {
	m, n := len(maze), len(maze[0])
	dir := []int{-1, 0, 1, 0, -1}
	queue := [][]int{entrance}
	maze[entrance[0]][entrance[1]] = '+'
	step := 0
	for len(queue) > 0 {
		nxt := [][]int{}
		for _, l := range queue {
			if (l[0] == 0 || l[0] == m-1 || l[1] == 0 || l[1] == n-1) && (l[0] != entrance[0] || l[1] != entrance[1]) {
				return step
			}
			for i := 0; i < 4; i++ {
				x, y := l[0]+dir[i], l[1]+dir[i+1]
				if x >= 0 && x < m && y >= 0 && y < n && maze[x][y] == '.' {
					nxt = append(nxt, []int{x, y})
					maze[x][y] = '+'
				}
			}
		}
		step++
		queue = nxt
	}
	return -1
}