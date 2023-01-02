func sumOfNumberAndReverse(num int) bool {
	for i := num / 2; i <= num; i++ {
		if i+reverse(i) == num {
			return true
		}
	}
	return false
}

func reverse(num int) int {
	r := 0
	for ; num > 0; num /= 10 {
		r = 10*r + num%10
	}
	return r
}