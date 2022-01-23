func maxPower(s string) int {
	cnt, tmp := 0, 0
	for i := range s {
		if i > 0 && s[i] == s[i-1] {
			tmp++
		} else {
			cnt = max(cnt, tmp)
			tmp = 1
		}
	}
	cnt = max(cnt, tmp)
	return cnt
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}