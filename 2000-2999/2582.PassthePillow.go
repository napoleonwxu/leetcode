func passThePillow(n int, time int) int {
	time %= 2 * (n - 1)
	if time < n {
		return time + 1
	}
	return 2*n - time - 1
}