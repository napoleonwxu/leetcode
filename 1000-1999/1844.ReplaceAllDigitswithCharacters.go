func replaceDigits(s string) string {
	chs := []byte(s)
	for i := 1; i < len(chs); i += 2 {
		chs[i] = chs[i-1] + chs[i] - '0'
	}
	return string(chs)
}