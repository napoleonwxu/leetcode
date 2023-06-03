const MOD = int(1e9 + 7)

func monkeyMove(n int) int {
	ans, base := 1, 2
	for n > 0 {
		if n%2 == 1 {
			ans = base * ans % MOD
		}
		base = base * base % MOD
		n /= 2
	}
	return (ans - 2 + MOD) % MOD
}