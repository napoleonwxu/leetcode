func minTimeToType(word string) int {
	ans := 0
	pre := byte('a')
	for _, ch := range []byte(word) {
		mi, ma := min(ch, pre), max(ch, pre)
		ans += int(min(ma-mi, 26+mi-ma)) + 1
		pre = ch
	}
	return ans
}

func max(x, y byte) byte {
	if x > y {
		return x
	}
	return y
}

func min(x, y byte) byte {
	if x < y {
		return x
	}
	return y
}