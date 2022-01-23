const MOD = 1e9 + 7

func numPrimeArrangements(n int) int {
    if n < 2 {
        return 1
    }
    cnt := cntPrime(n)
    ans := 1
    for i := 2; i <= cnt; i++ {
        ans = (ans*i) % MOD
    }
    for i := 2; i <= n-cnt; i++ {
        ans = (ans*i) % MOD
    }
    return ans
}

func cntPrime(n int) int {
    if n < 2 {
        return 0
    }
    prime := make([]bool, n+1)
    prime[2] = true
    for i := 3; i <= n; i += 2 {
        prime[i] = true
    }
    for i := 3; i*i <= n; i += 2 {
        if prime[i] {
            for j := i*i; j <= n; j += 2*i {
                prime[j] = false
            }
        }
    }
    cnt := 1
    for i := 3; i <= n; i += 2 {
        if prime[i] {
            cnt++
        }
    }
    return cnt
}