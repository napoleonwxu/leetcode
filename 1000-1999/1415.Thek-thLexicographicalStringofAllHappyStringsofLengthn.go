func getHappyString(n int, k int) string {
	perm := 1 << (n - 1)
	if k > 3*perm {
		return ""
	}
	ans := []byte{'a' + byte((k-1)/perm)}
	for perm > 1 {
		k -= (k - 1) / perm * perm
		perm >>= 1
		n := len(ans)
		if (k-1)/perm == 0 {
			if ans[n-1] == 'a' {
				ans = append(ans, 'b')
			} else {
				ans = append(ans, 'a')
			}
		} else {
			if ans[n-1] == 'c' {
				ans = append(ans, 'b')
			} else {
				ans = append(ans, 'c')
			}
		}
	}
	return string(ans)
}