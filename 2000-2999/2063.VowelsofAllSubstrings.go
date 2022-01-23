func countVowels(word string) int64 {
	ans := int64(0)
	n := len(word)
	chs := []byte(word)
	for i, ch := range chs {
		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
			ans += int64(i+1) * int64(n-i)
		}
	}
	return ans
}