func isPossible(n int, edges [][]int) bool {
	m := make(map[int]map[int]bool, n)
	for i := 1; i <= n; i++ {
		m[i] = make(map[int]bool)
	}
	for _, edge := range edges {
		m[edge[0]][edge[1]] = true
		m[edge[1]][edge[0]] = true
	}
	odds := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		if len(m[i])%2 == 1 {
			odds = append(odds, i)
		}
	}
	if len(odds) == 2 {
		for i := 1; i <= n; i++ {
			if !m[odds[0]][i] && !m[odds[1]][i] {
				return true
			}
		}
	}
	if len(odds) == 4 {
		return (!m[odds[0]][odds[1]] && !m[odds[2]][odds[3]]) ||
			(!m[odds[0]][odds[2]] && !m[odds[1]][odds[3]]) ||
			(!m[odds[0]][odds[3]] && !m[odds[1]][odds[2]])
	}
	return len(odds) == 0
}