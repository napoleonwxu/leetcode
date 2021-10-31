func maxMatrixSum(matrix [][]int) int64 {
	sum := int64(0)
	absMin, negCnt := int(1e5), 0
	for _, m := range matrix {
		for _, num := range m {
			if num < 0 {
				negCnt++
				num = -num
			}
			sum += int64(num)
			absMin = min(absMin, num)
		}
	}
	if negCnt%2 != 0 {
		return sum - 2*int64(absMin)
	}
	return sum
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}