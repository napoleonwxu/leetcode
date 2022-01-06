func countVowelSubstrings(word string) int {
	ans, vowels := 0, 0
	j, k := 0, 0
	cnt := make(map[byte]int)
	chs := []byte(word)
	for i, ch := range chs {
		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
			cnt[ch]++
			if cnt[ch] == 1 {
				vowels++
			}
			for ; vowels == 5; k++ {
				cnt[chs[k]]--
				if cnt[chs[k]] == 0 {
					vowels--
				}
			}
			ans += k - j
		} else {
			cnt['a'] = 0
			cnt['e'] = 0
			cnt['i'] = 0
			cnt['o'] = 0
			cnt['u'] = 0
			vowels = 0
			j = i + 1
			k = i + 1
		}
	}
	return ans
}