const MOD = 1e9 + 7

func countOrders(n int) int {
	cnt := 1
	// C(2n, 2) * C(2n-2, 2) * C(2n-4, 2) * ... * C(2, 2)
    for i := 2*n; i > 0; i -= 2 {
        cnt *= i*(i-1)/2
        cnt %= MOD
    }
    return cnt
}