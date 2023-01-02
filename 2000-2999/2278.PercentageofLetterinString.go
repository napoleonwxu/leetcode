func percentageLetter(s string, letter byte) int {
	cnt := 0
	for _, ch := range []byte(s) {
		if ch == letter {
			cnt++
		}
	}
	return 100 * cnt / len(s)
}