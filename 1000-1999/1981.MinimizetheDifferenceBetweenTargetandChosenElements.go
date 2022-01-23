func minimizeTheDifference(mat [][]int, target int) int {
	// O(target * n^2) + O(target)
	minSum := 0
	for _, row := range mat {
		mi := 70
		for _, num := range row {
			mi = min(mi, num)
		}
		minSum += mi
	}
	if minSum >= target {
		return minSum - target
	}

	m := make(map[int]bool)
	m[0] = true
	for _, row := range mat {
		tmp := make(map[int]bool)
		for _, num := range row {
			for sum := range m {
				sum += num
				if sum < 2*target-minSum {
					tmp[sum] = true
				}
			}
		}
		m = tmp
	}
	diff := 70 * 70
	for sum := range m {
		diff = min(diff, abs(sum-target))
	}
	return diff
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}