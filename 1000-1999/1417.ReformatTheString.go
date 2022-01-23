func reformat(s string) string {
	n := len(s)
	s1, s2 := []byte{}, []byte{}
	for i := 0; i < n; i++ {
		if s[i] >= '0' && s[i] <= '9' {
			s1 = append(s1, s[i])
		} else {
			s2 = append(s2, s[i])
		}
	}
	if len(s1) < len(s2) {
		s1, s2 = s2, s1
	}
	if len(s1)-len(s2) > 1 {
		return ""
	}
	ans := make([]byte, n)
	for i := 0; i < len(s2); i++ {
		ans[2*i] = s1[i]
		ans[2*i+1] = s2[i]
	}
	if len(s1) > len(s2) {
		ans[n-1] = s1[len(s1)-1]
	}
	return string(ans)
}
