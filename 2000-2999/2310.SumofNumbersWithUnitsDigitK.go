func minimumNumbers(num int, k int) int {
	if num == 0 {
		return 0
	}
	for i := 1; i <= 10 && i*k <= num; i++ {
		if i*k%10 == num%10 {
			return i
		}
	}
	return -1
}