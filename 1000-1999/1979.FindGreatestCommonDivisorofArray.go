func findGCD(nums []int) int {
	mi, ma := 1000, 0
	for _, num := range nums {
		mi = min(mi, num)
		ma = max(ma, num)
	}
	return gcd(mi, ma)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}