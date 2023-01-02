func greatestLetter(s string) string {
	chs := []byte(s)
	chMap := make(map[byte]bool, 26*2)
	for _, ch := range chs {
		chMap[ch] = true
	}
	for ch := byte('Z'); ch >= 'A'; ch-- {
		if chMap[ch] && chMap[ch+'a'-'A'] {
			return string(ch)
		}
	}
	return ""
}