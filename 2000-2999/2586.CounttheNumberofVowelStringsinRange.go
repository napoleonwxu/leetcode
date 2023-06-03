func vowelStrings(words []string, left int, right int) int {
	cnt := 0
	for i := left; i <= right; i++ {
		n := len(words[i])
		if isVowel(words[i][0]) && isVowel(words[i][n-1]) {
			cnt++
		}
	}
	return cnt
}

func isVowel(ch byte) bool {
	return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u'
}