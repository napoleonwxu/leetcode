func truncateSentence(s string, k int) string {
	//return strings.Join(strings.Split(s, " ")[:k], " ")
	for i, ch := range s {
		if ch == ' ' {
			k--
			if k == 0 {
				return s[:i]
			}
		}
	}
	return s
}