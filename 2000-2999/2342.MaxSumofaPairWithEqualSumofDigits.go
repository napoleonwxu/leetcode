func maximumSum(nums []int) int {
	ans := -1
	m := make(map[int]int)
	for _, num := range nums {
		sum := 0
		for tmp := num; tmp > 0; tmp /= 10 {
			sum += tmp % 10
		}
		if pre, ok := m[sum]; ok {
			ans = max(ans, num+pre)
		}
		m[sum] = max(m[sum], num)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}