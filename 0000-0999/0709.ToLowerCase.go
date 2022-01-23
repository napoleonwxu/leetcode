func toLowerCase(s string) string {
	//return strings.ToLower(s)
	chs := []byte(s)
	for i, ch := range chs {
		if ch >= 'A' && ch <= 'Z' {
			chs[i] += 32
		}
	}
	return string(chs)
}