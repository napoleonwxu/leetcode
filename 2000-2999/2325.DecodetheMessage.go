func decodeMessage(key string, message string) string {
	m := make(map[byte]byte, 27)
	m[' '] = ' '
	cur := byte('a')
	for _, ch := range []byte(key) {
		if _, ok := m[ch]; !ok {
			m[ch] = cur
			cur++
		}
	}
	ans := make([]byte, len(message))
	for i, ch := range []byte(message) {
		ans[i] = m[ch]
	}
	return string(ans)
}