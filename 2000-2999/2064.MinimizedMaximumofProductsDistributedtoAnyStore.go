func minimizedMaximum(n int, quantities []int) int {
	left, right := 1, 0
	for _, q := range quantities {
		right = max(right, q)
	}
	for left < right {
		mid, cnt := (left+right)/2, 0
		for _, q := range quantities {
			cnt += (q + mid - 1) / mid
		}
		if cnt <= n {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}