import "math"

func countTriples(n int) int {
	cnt := 0
	for a := 1; a <= n; a++ {
		for b := a; b <= n; b++ {
			sq := a*a + b*b
			c := int(math.Sqrt(float64(sq)))
			if c*c == sq && c <= n {
				cnt += 2
			}
		}
	}
	return cnt
}