func isThree(n int) bool {
	cnt, d := 0, 1
	for ; d*d < n; d++ {
		if n%d == 0 {
			cnt += 2
			if cnt > 3 {
				return false
			}
		}
	}
	if d*d == n {
		cnt++
	}
	return cnt == 3
}