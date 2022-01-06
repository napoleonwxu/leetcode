import "strings"

func shortestCompletingWord(licensePlate string, words []string) string {
	lic := strings.ToLower(licensePlate)
	licCnt := count(lic)
	ans := ""
	for _, word := range words {
		if (ans == "" || len(word) < len(ans)) && complete(licCnt, count(word)) {
			ans = word
		}
	}
	return ans
}

func count(s string) []int {
	cnt := make([]int, 26)
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			cnt[ch-'a']++
		}
	}
	return cnt
}

func complete(licCnt, wordCnt []int) bool {
	for i := 0; i < len(licCnt); i++ {
		if wordCnt[i] < licCnt[i] {
			return false
		}
	}
	return true
}