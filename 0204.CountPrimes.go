func countPrimes(n int) int {
    if n <= 2 {
        return 0
    }
    prime := make([]bool, n)
    prime[2] = true
    for i := 3; i < n; i += 2 {
        prime[i] = true
    }
    for i := 3; i*i < n; i += 2 {
        if prime[i] {
            for j := i*i; j < n; j += i+i {
                prime[j] = false
            }
        }
    }
    cnt := 1
    for i := 3; i < n; i += 2 {
        if prime[i] {
            cnt++
        }
    }
    return cnt
}