func firstPalindrome(words []string) string {
	for _, word := range words {
		if isPalindromic(word) {
			return word
		}
	}
	return ""
}

func isPalindromic(word string) bool {
	n := len(word)
	for i := 0; i < n/2; i++ {
		if word[i] != word[n-1-i] {
			return false
		}
	}
	return true
}