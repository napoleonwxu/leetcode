import "strings"

func countValidWords(sentence string) int {
	words := strings.Fields(sentence)
	cnt := 0
	for _, word := range words {
		if validToken(word) {
			cnt++
		}
	}
	return cnt
}

func validToken(word string) bool {
	bytes := []byte(word)
	n, hyphen := len(bytes), 0
	for i, ch := range bytes {
		if ch >= '0' && ch <= '9' {
			return false
		}
		if ch == '-' {
			if hyphen >= 1 || i <= 0 || i >= n-1 || !isLower(bytes[i-1]) || !isLower(bytes[i+1]) {
				return false
			}
			hyphen++
		}
		if (ch == '!' || ch == '.' || ch == ',') && i != n-1 {
			return false
		}
	}
	return true
}

func isLower(ch byte) bool {
	return ch >= 'a' && ch <= 'z'
}
