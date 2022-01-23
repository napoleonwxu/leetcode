func minimumRefill(plants []int, capacityA int, capacityB int) int {
	left, right := 0, len(plants)-1
	a, b := capacityA, capacityB
	cnt := 0
	for ; left < right; left, right = left+1, right-1 {
		if a < plants[left] {
			cnt++
			a = capacityA
		}
		a -= plants[left]

		if b < plants[right] {
			cnt++
			b = capacityB
		}
		b -= plants[right]
	}
	if left == right {
		if max(a, b) < plants[left] {
			cnt++
		}
	}
	return cnt
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}