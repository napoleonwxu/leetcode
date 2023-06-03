func removeTrailingZeros(num string) string {
	//return strings.TrimRight(num, "0")
	chs := []byte(num)
	i := len(chs) - 1
	for ; i >= 0 && chs[i] == '0'; i-- {
	}
	return string(chs[:i+1])
}