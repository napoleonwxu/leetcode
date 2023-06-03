func primeSubOperation(nums []int) bool {
	primes := sieve(1000)
	n := len(nums)
	pre := nums[n-1]
	for i := n - 2; i >= 0; i-- {
		if nums[i] < pre {
			pre = nums[i]
			continue
		}
		found := false
		for _, prime := range primes {
			if prime >= nums[i] {
				return false
			}
			if nums[i]-prime < pre {
				pre = nums[i] - prime
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func sieve(n int) []int {
	isPrimes := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrimes[i] = true
	}
	for i := 2; i*i <= n; i++ {
		if isPrimes[i] {
			for j := i * i; j <= n; j += i {
				isPrimes[j] = false
			}
		}
	}
	primes := []int{}
	for i, isPrime := range isPrimes {
		if isPrime {
			primes = append(primes, i)
		}
	}
	return primes
}