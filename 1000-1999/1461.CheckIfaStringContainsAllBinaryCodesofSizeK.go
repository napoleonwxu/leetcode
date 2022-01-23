func hasAllCodes(s string, k int) bool {
	set := make(map[string]bool)
	for i := 0; i <= len(s)-k; i++ {
		set[s[i:i+k]] = true
	}
	return len(set) == 1<<k
}