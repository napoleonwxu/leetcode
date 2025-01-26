func nonSpecialCount(l int, r int) int {
    ma := int(math.Sqrt(float64(r)))
    primes := make([]bool, ma+1)
    if ma >= 2 {
        primes[2] = true
    }
    for i := 3; i <= ma; i += 2 {
        primes[i] = true
    }
    for i := 3; i <= ma; i += 2 {
        if primes[i] {
            for j := i; i*j <= ma; j += 2 {
                primes[i*j] = false
            }
        }
    }
    cnt := 0
    for i := 2; i <= ma; i++ {
        if primes[i] && i*i >= l && i*i <= r {
            cnt++
        }
    }
    return r-l+1-cnt
}
