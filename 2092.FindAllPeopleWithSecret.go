import "sort"

func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
	// BFS
	sort.Slice(meetings, func(i, j int) bool {
		if meetings[i][2] < meetings[j][2] {
			return true
		}
		return false
	})
	m := make(map[int]bool)
	m[0] = true
	m[firstPerson] = true
	idx := 0
	for idx < len(meetings) {
		graph := make(map[int][]int)
		tmp := make(map[int]bool)
		i := idx
		for ; i < len(meetings) && meetings[i][2] == meetings[idx][2]; i++ {
			graph[meetings[i][0]] = append(graph[meetings[i][0]], meetings[i][1])
			graph[meetings[i][1]] = append(graph[meetings[i][1]], meetings[i][0])
			if m[meetings[i][0]] {
				tmp[meetings[i][0]] = true
			}
			if m[meetings[i][1]] {
				tmp[meetings[i][1]] = true
			}
		}
		queue := make([]int, 0, len(tmp))
		for p := range tmp {
			queue = append(queue, p)
		}
		for len(queue) > 0 {
			p := queue[0]
			queue = queue[1:]
			for _, nxt := range graph[p] {
				if !m[nxt] {
					m[nxt] = true
					queue = append(queue, nxt)
				}
			}
		}
		idx = i
	}
	ans := make([]int, 0, len(m))
	for person := range m {
		ans = append(ans, person)
	}
	return ans
}