func findValidSplit(nums []int) int {
	fi := make(map[int]int)
	line := make([]int, 10001)
	for i, num := range nums {
		for _, f := range getFactors(num) {
			if _, ok := fi[f]; !ok {
				fi[f] = i
			}
			line[fi[f]]++
			line[i]--
		}
	}
	for i := 0; i < len(line)-1; i++ {
		if line[i] == 0 && i < len(nums)-1 {
			return i
		}
		line[i+1] += line[i]
	}
	return -1
}

func getFactors(num int) []int {
	factors := []int{}
	for f := 2; num > 1 && f < 1000; f += 1 + f%2 {
		if num%f == 0 {
			factors = append(factors, f)
		}
		for num%f == 0 {
			num /= f
		}
	}
	if num > 1 {
		factors = append(factors, num)
	}
	return factors
}