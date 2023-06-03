func diagonalPrime(nums [][]int) int {
	ans, n := 0, len(nums)
	for i := 0; i < n; i++ {
		if nums[i][i] > ans && isPrime(nums[i][i]) {
			ans = nums[i][i]
		}
		if nums[i][n-1-i] > ans && isPrime(nums[i][n-1-i]) {
			ans = nums[i][n-1-i]
		}
	}
	return ans
}

func isPrime(num int) bool {
	if num <= 2 || num%2 == 0 {
		return num == 2
	}
	for i := 3; i*i <= num; i += 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}