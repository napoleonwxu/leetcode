func repeatedCharacter(s string) byte {
	chs := make([]bool, 26)
	for _, ch := range []byte(s) {
		if chs[ch-'a'] {
			return ch
		}
		chs[ch-'a'] = true
	}
	return 'A'
}