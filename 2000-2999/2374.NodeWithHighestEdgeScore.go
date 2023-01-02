func edgeScore(edges []int) int {
	scores := make([]int, len(edges))
	for i, node := range edges {
		scores[node] += i
	}
	idx, max := -1, -1
	for node, score := range scores {
		if score > max || (score == max && node < idx) {
			idx, max = node, score
		}
	}
	return idx
}