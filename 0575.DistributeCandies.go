func distributeCandies(candyType []int) int {
	m := make(map[int]bool)
	for _, candy := range candyType {
		m[candy] = true
	}
	return min(len(m), len(candyType)/2)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}