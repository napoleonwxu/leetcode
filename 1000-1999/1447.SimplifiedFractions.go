import "strconv"

func simplifiedFractions(n int) []string {
	ans := []string{}
	for i := 1; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			if gcd(i, j) == 1 {
				ans = append(ans, strconv.Itoa(i)+"/"+strconv.Itoa(j))
			}
		}
	}
	return ans
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}