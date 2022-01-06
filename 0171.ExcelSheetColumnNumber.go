func titleToNumber(columnTitle string) int {
	ans := 0
	for _, ch := range columnTitle {
		ans = 26*ans + int(ch-'A'+1)
	}
	return ans
}