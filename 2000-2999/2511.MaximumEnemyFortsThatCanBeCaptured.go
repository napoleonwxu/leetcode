func captureForts(forts []int) int {
	l, ans := 0, 0
	for r := 0; r < len(forts); r++ {
		if forts[r] != 0 {
			if forts[l]+forts[r] == 0 {
				ans = max(ans, r-l-1)
			}
			l = r
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}