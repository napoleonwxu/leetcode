import "strings"

func numOfStrings(patterns []string, word string) int {
	cnt := 0
	for _, pattern := range patterns {
		if strings.Contains(word, pattern) {
			cnt++
		}
	}
	return cnt
}