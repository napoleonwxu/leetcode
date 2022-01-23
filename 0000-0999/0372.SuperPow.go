const MOD = 1337

func superPow(a int, b []int) int {
	n := len(b)
	if n == 0 {
		return 1
	}
	d := b[n-1]
	b = b[:n-1]
	return power(superPow(a, b), 10) * power(a, d) % MOD
}

func power(a, d int) int {
	a %= MOD
	p := 1
	for d > 0 {
		p = p * a % MOD
		d--
	}
	return p
}