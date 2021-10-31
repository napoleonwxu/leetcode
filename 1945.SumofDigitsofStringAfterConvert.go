func getLucky(s string, k int) int {
	num := 0
	for _, ch := range []byte(s) {
		num += transform(int(ch-'a') + 1)
	}
	for num >= 10 && k > 1 {
		num = transform(num)
		k--
	}
	return num
}

func transform(num int) int {
	ans := 0
	for num > 0 {
		ans += num % 10
		num /= 10
	}
	return ans
}
