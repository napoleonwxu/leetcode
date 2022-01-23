func isPerfectSquare(num int) bool {
	left, right := 1, num
	for left <= right {
		mid := (left + right) / 2
		if mid*mid == num {
			return true
		}
		if mid*mid < num {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}