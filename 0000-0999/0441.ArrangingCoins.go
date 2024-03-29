func arrangeCoins(n int) int {
	left, right := 1, (n+1)/2
	for left <= right {
		mid := (left + right) / 2
		if (1+mid)*mid/2 <= n {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return right
}