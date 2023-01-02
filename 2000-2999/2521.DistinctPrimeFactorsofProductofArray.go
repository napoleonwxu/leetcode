func distinctPrimeFactors(nums []int) int {
	primes := make(map[int]bool)
	for _, num := range nums {
		div := 2
		for num > 1 {
			if num%div == 0 {
				primes[div] = true
				num /= div
			} else {
				div++
			}
		}
	}
	return len(primes)
}