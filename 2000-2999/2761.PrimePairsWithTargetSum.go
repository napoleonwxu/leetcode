func findPrimePairs(n int) [][]int {
    primes := calPrimes(n)
    ans := [][]int{}
    for i := 2; 2*i <= n; i++ {
        if primes[i] && primes[n-i] {
            ans = append(ans, []int{i, n - i})
        }
    }
    return ans
}

func calPrimes(n int) []bool {
    primes := make([]bool, n)
    if n <= 2 {
        return primes
    }
    primes[2] = true
    for i := 3; i < n; i += 2 {
        primes[i] = true
    }
    for i := 3; i*i < n; i += 2 {
        for j := i * i; j < n; j += i {
            primes[j] = false
        }
    }
    return primes
}