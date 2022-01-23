func makeFancyString(s string) string {
	if len(s) <= 2 {
		return s
	}
	chs := []byte(s)
	i := 2
	for j := 2; j < len(chs); j++ {
		if chs[j] != chs[i-1] || chs[j] != chs[i-2] {
			chs[i] = chs[j]
			i++
		}
	}
	return string(chs[:i])
}