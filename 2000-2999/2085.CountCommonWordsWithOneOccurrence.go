func countWords(words1 []string, words2 []string) int {
	m := make(map[string]int)
	for _, word := range words1 {
		m[word]++
	}
	for _, word := range words2 {
		if m[word] <= 1 {
			m[word]--
		}
	}
	ans := 0
	for _, cnt := range m {
		if cnt == 0 {
			ans++
		}
	}
	return ans
}