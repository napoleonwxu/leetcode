func maximumNumber(num string, change []int) string {
	mutated := false
	nums := []byte(num)
	for i, ch := range nums {
		tmp := byte(change[int(ch-'0')]) + '0'
		if tmp >= ch {
			nums[i] = tmp
			if tmp > ch {
				mutated = true
			}
		} else if mutated {
			break
		}
	}
	return string(nums)
}