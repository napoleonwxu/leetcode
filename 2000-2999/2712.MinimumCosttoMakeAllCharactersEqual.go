func minimumCost(s string) int64 {
	n := len(s)
	ans := int64(0)
	for i := 1; i < n; i++ {
		if s[i] != s[i-1] {
			ans += int64(min(i, n-i))
		}
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}