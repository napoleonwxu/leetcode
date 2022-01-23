func numberOfWeeks(milestones []int) int64 {
	var sum int64
	mx := 0
	for _, m := range milestones {
		sum += int64(m)
		mx = max(mx, m)
	}
	return min(sum, 2*(sum-int64(mx))+1)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}