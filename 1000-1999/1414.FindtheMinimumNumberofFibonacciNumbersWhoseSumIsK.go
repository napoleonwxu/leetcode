func findMinFibonacciNumbers(k int) int {
	fib := []int{1, 1}
	for fib[len(fib)-1] < k {
		n := len(fib)
		fib = append(fib, fib[n-2]+fib[n-1])
	}
	cnt := 0
	for i := len(fib) - 1; k > 0 && i > 0; i-- {
		if k >= fib[i] {
			k -= fib[i]
			cnt++
		}
	}
	return cnt
}