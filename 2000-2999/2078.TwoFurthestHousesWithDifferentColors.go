func maxDistance(colors []int) int {
	n := len(colors)
	i, j := 0, n-1
	for ; colors[i] == colors[n-1]; i++ {
	}
	for ; colors[j] == colors[0]; j-- {
	}
	return max(j, n-i-1)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}