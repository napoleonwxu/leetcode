func countDigits(num int) int {
	cnt := 0
	for tmp := num; tmp > 0; tmp /= 10 {
		if num%(tmp%10) == 0 {
			cnt++
		}
	}
	return cnt
}