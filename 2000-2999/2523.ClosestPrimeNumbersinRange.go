import "math"

func closestPrimes(left int, right int) []int {
    primes := []int{}
    if left <= 2 {
        primes = append(primes, 2)
    }
    for num := max(3, 2*(left/2)+1); num <= right; num += 2 {
        if isPrime(num) {
            n := len(primes)
            if n > 0 && num-primes[n-1] <= 2 { // [2, 3] or twin primes
                return []int{primes[n-1], num}
            }
            primes = append(primes, num)
        }
    }
    if len(primes) < 2 {
        return []int{-1, -1}
    }
    ans := primes[:2]
    for i := 1; i < len(primes); i++ {
        if primes[i]-primes[i-1] < ans[1]-ans[0] {
            ans = primes[i-1 : i+1]
        }
    }
    return ans
}

func isPrime(num int) bool {
    if num <= 1 {
        return false
    }
    if num <= 3 {
        return true
    }
    maxDiv := int(math.Sqrt(float64(num)))
    for div := 3; div <= maxDiv; div += 2 {
        if num%div == 0 {
            return false
        }
    }
    return true
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}