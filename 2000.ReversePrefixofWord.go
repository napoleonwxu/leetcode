func reversePrefix(word string, ch byte) string {
	ans := []byte(word)
	for i, w := range ans {
		if w == ch {
			reverse(ans, i)
			break
		}
	}
	return string(ans)
}

func reverse(bytes []byte, idx int) {
	for i := 0; i <= idx/2; i++ {
		bytes[i], bytes[idx-i] = bytes[idx-i], bytes[i]
	}
}