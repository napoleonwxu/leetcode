func countPoints(rings string) int {
	colors := make([]int, 10)
	for i := 0; i < len(rings); i += 2 {
		color := 1
		if rings[i] == 'G' {
			color = 2
		} else if rings[i] == 'B' {
			color = 4
		}
		colors[rings[i+1]-'0'] |= color
	}
	cnt := 0
	for _, color := range colors {
		if color == 7 {
			cnt++
		}
	}
	return cnt
}