func removeAnagrams(words []string) []string {
	n := len(words)
	cnts := make([][26]int, n)
	for i, word := range words {
		for _, ch := range word {
			cnts[i][ch-'a']++
		}
	}
	ans := make([]string, 0, n)
	for i, word := range words {
		if i == 0 || !anagram(cnts[i], cnts[i-1]) {
			ans = append(ans, word)
		}
	}
	return ans
}

func anagram(cnt1, cnt2 [26]int) bool {
	for i := 0; i < len(cnt1); i++ {
		if cnt1[i] != cnt2[i] {
			return false
		}
	}
	return true
}